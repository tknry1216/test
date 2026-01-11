package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/takenariyamamoto/golang/internal/dbmgr"
)

var (
	dbURL string
	db    *dbmgr.DBManager
)

var rootCmd = &cobra.Command{
	Use:   "golang-cli",
	Short: "PostgreSQL CRUD操作のためのCLIツール",
	Long:  `Go 1.25とCobraで構築されたPostgreSQL操作用CLIツール`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// データベース接続
		var err error
		db, err = dbmgr.New(context.Background(), dbURL)
		if err != nil {
			return fmt.Errorf("データベース接続に失敗しました: %w", err)
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		// データベース接続を閉じる
		if db != nil {
			db.Close()
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// グローバルフラグ
	rootCmd.PersistentFlags().StringVar(&dbURL, "db-url", getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"), "データベース接続URL")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
