package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/takenariyamamoto/golang/internal/dbmgr"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "ユーザー管理コマンド",
	Long:  `ユーザーの一覧表示、取得、作成などの操作を行います`,
}

var usersListCmd = &cobra.Command{
	Use:   "list",
	Short: "全ユーザーを一覧表示",
	Long:  `データベース内の全ユーザーを取得して表示します`,
	RunE: func(cmd *cobra.Command, args []string) error {
		users, err := db.GetAllUsers(context.Background())
		if err != nil {
			return fmt.Errorf("ユーザー取得に失敗しました: %w", err)
		}

		// JSON形式で出力
		output, err := json.MarshalIndent(users, "", "  ")
		if err != nil {
			return fmt.Errorf("JSON変換に失敗しました: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var usersGetCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "指定したIDのユーザーを取得",
	Long:  `IDを指定してユーザー情報を取得します`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("無効なID: %s", args[0])
		}

		user, err := db.GetUserByID(context.Background(), id)
		if err != nil {
			return fmt.Errorf("ユーザー取得に失敗しました: %w", err)
		}

		// JSON形式で出力
		output, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			return fmt.Errorf("JSON変換に失敗しました: %w", err)
		}

		fmt.Println(string(output))
		return nil
	},
}

var (
	userName  string
	userEmail string
)

var usersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "新しいユーザーを作成",
	Long:  `名前とメールアドレスを指定して新しいユーザーを作成します`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userName == "" || userEmail == "" {
			return fmt.Errorf("--name と --email は必須です")
		}

		params := dbmgr.CreateUserParams{
			Name:  userName,
			Email: userEmail,
		}

		user, err := db.CreateUser(context.Background(), params)
		if err != nil {
			return fmt.Errorf("ユーザー作成に失敗しました: %w", err)
		}

		// JSON形式で出力
		output, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			return fmt.Errorf("JSON変換に失敗しました: %w", err)
		}

		fmt.Println("✅ ユーザーを作成しました:")
		fmt.Println(string(output))
		return nil
	},
}

func init() {
	// usersコマンドをrootに追加
	rootCmd.AddCommand(usersCmd)

	// サブコマンドをusersに追加
	usersCmd.AddCommand(usersListCmd)
	usersCmd.AddCommand(usersGetCmd)
	usersCmd.AddCommand(usersCreateCmd)

	// createコマンドのフラグ
	usersCreateCmd.Flags().StringVarP(&userName, "name", "n", "", "ユーザー名（必須）")
	usersCreateCmd.Flags().StringVarP(&userEmail, "email", "e", "", "メールアドレス（必須）")
	usersCreateCmd.MarkFlagRequired("name")
	usersCreateCmd.MarkFlagRequired("email")
}

