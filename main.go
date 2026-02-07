package main

import (
	"html/template"
	"log"
	"os"

	"perennial-wisdom/db"
	"perennial-wisdom/router"
	"perennial-wisdom/store"
)

func main() {
	// Database path — configurable via env, defaults to local file
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "wisdom.db"
	}

	// Open SQLite (pure Go, embedded, no CGO)
	database := db.Open(dbPath)
	defer database.Close()

	// Run migrations + seed data
	db.Migrate(database)
	db.Seed(database)

	// Create query layer
	queries := db.NewQueries(database)

	// Parse HTML templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	template.Must(tmpl.ParseGlob("templates/partials/*.html"))

	// In-memory store still available for JSON API
	s := store.New()

	// Wire all routes with explicit dependencies
	r := router.Setup(s, queries, tmpl)

	// Port — configurable via env, defaults to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Perennial Wisdom API starting on :%s", port)
	r.Run(":" + port)
}
