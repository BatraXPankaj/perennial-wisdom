package db

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

// Open opens (or creates) a SQLite database at the given path.
// Uses WAL mode for better concurrent read performance.
// Returns a sqlx.DB which wraps database/sql with named-query support.
func Open(dbPath string) *sqlx.DB {
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("failed to open database at %s: %v", dbPath, err)
	}

	// WAL mode: concurrent readers, single writer, no locks on reads
	conn.Exec("PRAGMA journal_mode=WAL")
	conn.Exec("PRAGMA foreign_keys=ON")

	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return sqlx.NewDb(conn, "sqlite")
}
