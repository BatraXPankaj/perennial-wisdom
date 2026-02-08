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
	// Open PostgreSQL connection
	database := db.Open("")
	defer database.Close()

	// Create query layer
	queries := db.NewQueries(database)

	// Parse HTML templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	template.Must(tmpl.ParseGlob("templates/partials/*.html"))

	// In-memory store still available for JSON API (legacy, can be removed later)
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
