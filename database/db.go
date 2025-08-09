package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sqlx.DB

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Environment variable %s belum di-set", key)
	}
	return value
}

func Connect() {
	var err error

	// Ambil environment variables dengan pengecekan
	user := getEnv("DB_USER")
	password := getEnv("DB_PASSWORD")
	host := getEnv("DB_HOST")
	port := getEnv("DB_PORT")
	name := getEnv("DB_NAME")

	// Susun DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		name,
	)

	// Koneksi ke database
	DB, err = sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("❌ Gagal connect ke database: %v", err)
	}

	fmt.Println("✅ Database connected")
}
