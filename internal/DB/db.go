package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect to PostgreSQL

func OpenDB() *sql.DB {
	// Connection string
	dsn := "postgresql://neondb_owner:npg_aQvD60wWULIo@ep-small-salad-ad8vc158-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect:", err)
	}

	return db
}