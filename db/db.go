package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Open connects to PostgreSQL using DATABASE_URL or constructs from components.
// Returns a sqlx.DB which wraps database/sql with named-query support.
func Open(dbPath string) *sqlx.DB {
	// Check for PostgreSQL connection string
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Construct from components or use default
		host := os.Getenv("DB_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}
		user := os.Getenv("DB_USER")
		if user == "" {
			user = "wisdom"
		}
		password := os.Getenv("DB_PASSWORD")
		if password == "" {
			password = "perennial2026"
		}
		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			dbname = "perennial_wisdom"
		}
		sslmode := os.Getenv("DB_SSLMODE")
		if sslmode == "" {
			sslmode = "disable"
		}
		connStr = "host=" + host + " port=" + port + " user=" + user +
			" password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return sqlx.NewDb(conn, "postgres")
}
