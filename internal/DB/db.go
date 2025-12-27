package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect to PostgreSQL

func OpenDB() *sql.DB {
	// Connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect:", err)
	}
	fmt.Println("Successfully connected to PostgreSQL database!")
	return db
}