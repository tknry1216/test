package dbmgr

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager struct {
	pool *pgxpool.Pool
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserParams struct {
	Name  string
	Email string
}

type UpdateUserParams struct {
	Name  *string
	Email *string
}

// New creates a new DBManager instance
func New(ctx context.Context, dbURL string) (*DBManager, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// 接続確認
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DBManager{pool: pool}, nil
}

// Close closes the database connection pool
func (db *DBManager) Close() {
	db.pool.Close()
}

// GetAllUsers retrieves all users from the database
func (db *DBManager) GetAllUsers(ctx context.Context) ([]User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users ORDER BY id"
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID
func (db *DBManager) GetUserByID(ctx context.Context, id int) (*User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"

	var user User
	err := db.pool.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// CreateUser creates a new user
func (db *DBManager) CreateUser(ctx context.Context, params CreateUserParams) (*User, error) {
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, name, email, created_at, updated_at
	`

	var user User
	err := db.pool.QueryRow(ctx, query, params.Name, params.Email).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

// UpdateUser updates a user
func (db *DBManager) UpdateUser(ctx context.Context, id int, params UpdateUserParams) (*User, error) {
	query := "UPDATE users SET updated_at = CURRENT_TIMESTAMP"
	args := []interface{}{}
	argCount := 1

	if params.Name != nil {
		query += fmt.Sprintf(", name = $%d", argCount)
		args = append(args, *params.Name)
		argCount++
	}
	if params.Email != nil {
		query += fmt.Sprintf(", email = $%d", argCount)
		args = append(args, *params.Email)
		argCount++
	}

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, name, email, created_at, updated_at", argCount)
	args = append(args, id)

	var user User
	err := db.pool.QueryRow(ctx, query, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}

// DeleteUser deletes a user by ID
func (db *DBManager) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := db.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
