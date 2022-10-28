package config

import (
	"fmt"
	"os"
)

// DSNはdataSourceNameを返します、もし必須の環境変数が設定されてなかった場合はerrorを返します
func DSN() (string, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbDatabase == "" {
		return "", fmt.Errorf("ERROR : required environment variable not found")
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
	) + "parseTime=true&collation=utf8mb4_bin", nil
}
