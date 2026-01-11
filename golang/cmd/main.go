package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/takenariyamamoto/golang/internal/dbmgr"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

var db *dbmgr.DBManager

func main() {
	// データベース接続
	dbURL := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable")

	var err error
	db, err = dbmgr.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("データベース接続に失敗しました: %v", err)
	}
	defer db.Close()

	log.Println("データベースに接続しました")

	// ルーター設定
	mux := http.NewServeMux()

	// ヘルスチェック
	mux.HandleFunc("GET /health", healthHandler)

	// ユーザーAPI
	mux.HandleFunc("GET /api/v1/users", getUsersHandler)
	mux.HandleFunc("GET /api/v1/users/{id}", getUserHandler)
	mux.HandleFunc("POST /api/v1/users", createUserHandler)
	mux.HandleFunc("PUT /api/v1/users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /api/v1/users/{id}", deleteUserHandler)

	// サーバー起動
	port := getEnv("PORT", "8080")
	log.Printf("サーバーをポート %s で起動します", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("サーバー起動に失敗しました: %v", err)
	}
}

// GET /health - ヘルスチェック
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GET /api/v1/users - 全ユーザー取得
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "ユーザー取得に失敗しました")
		return
	}

	writeJSON(w, http.StatusOK, users)
}

// GET /api/v1/users/{id} - 特定ユーザー取得
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "無効なIDです")
		return
	}

	user, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			writeError(w, http.StatusNotFound, "ユーザーが見つかりません")
			return
		}
		writeError(w, http.StatusInternalServerError, "ユーザー取得に失敗しました")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// POST /api/v1/users - ユーザー作成
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "リクエストの解析に失敗しました")
		return
	}

	if req.Name == "" || req.Email == "" {
		writeError(w, http.StatusBadRequest, "nameとemailは必須です")
		return
	}

	params := dbmgr.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := db.CreateUser(r.Context(), params)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "ユーザー作成に失敗しました")
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

// PUT /api/v1/users/{id} - ユーザー更新
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "無効なIDです")
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "リクエストの解析に失敗しました")
		return
	}

	params := dbmgr.UpdateUserParams{}
	if req.Name != "" {
		params.Name = &req.Name
	}
	if req.Email != "" {
		params.Email = &req.Email
	}

	user, err := db.UpdateUser(r.Context(), id, params)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			writeError(w, http.StatusNotFound, "ユーザーが見つかりません")
			return
		}
		writeError(w, http.StatusInternalServerError, "ユーザー更新に失敗しました")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// DELETE /api/v1/users/{id} - ユーザー削除
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "無効なIDです")
		return
	}

	if err := db.DeleteUser(r.Context(), id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			writeError(w, http.StatusNotFound, "ユーザーが見つかりません")
			return
		}
		writeError(w, http.StatusInternalServerError, "ユーザー削除に失敗しました")
		return
	}

	writeJSON(w, http.StatusOK, MessageResponse{Message: "ユーザーを削除しました"})
}

// Helper functions

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
