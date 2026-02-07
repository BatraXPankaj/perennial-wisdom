package db

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"

	"perennial-wisdom/store"
)

// Seed populates the database with perennial wisdom data.
// Skips if data already exists (idempotent).
func Seed(db *sqlx.DB) {
	var count int
	db.Get(&count, "SELECT COUNT(*) FROM quotes")
	if count > 0 {
		log.Printf("database already seeded (%d quotes), skipping", count)
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatalf("failed to begin seed transaction: %v", err)
	}
	defer tx.Rollback()

	// Philosophies
	for _, p := range store.SeedPhilosophies() {
		principles, _ := json.Marshal(p.CorePrinciples)
		tx.Exec("INSERT INTO philosophies (id, name, origin, core_principles) VALUES (?, ?, ?, ?)",
			p.ID, p.Name, p.Origin, string(principles))
		for _, rid := range p.RelatedIDs {
			tx.Exec("INSERT OR IGNORE INTO philosophy_relations (philosophy_id, related_id) VALUES (?, ?)",
				p.ID, rid)
		}
	}

	// Philosophers
	for _, p := range store.SeedPhilosophers() {
		teachings, _ := json.Marshal(p.KeyTeachings)
		tx.Exec("INSERT INTO philosophers (id, name, philosophy_id, era, bio, key_teachings) VALUES (?, ?, ?, ?, ?, ?)",
			p.ID, p.Name, p.PhilosophyID, p.Era, p.Bio, string(teachings))
	}

	// Themes
	for _, t := range store.SeedThemes() {
		tx.Exec("INSERT INTO themes (id, name, description) VALUES (?, ?, ?)",
			t.ID, t.Name, t.Description)
		for _, pid := range t.PhilosophyIDs {
			tx.Exec("INSERT OR IGNORE INTO theme_philosophies (theme_id, philosophy_id) VALUES (?, ?)",
				t.ID, pid)
		}
	}

	// Evidence
	for _, e := range store.SeedEvidence() {
		tx.Exec("INSERT INTO evidence (id, title, finding, field, source) VALUES (?, ?, ?, ?, ?)",
			e.ID, e.Title, e.Finding, e.Field, e.Source)
		for _, tid := range e.ThemeIDs {
			tx.Exec("INSERT OR IGNORE INTO evidence_themes (evidence_id, theme_id) VALUES (?, ?)",
				e.ID, tid)
		}
	}

	// Quotes
	for _, q := range store.SeedQuotes() {
		tx.Exec("INSERT INTO quotes (id, text, philosopher_id, philosophy_id, source) VALUES (?, ?, ?, ?, ?)",
			q.ID, q.Text, q.PhilosopherID, q.PhilosophyID, q.Source)
		for _, tid := range q.ThemeIDs {
			tx.Exec("INSERT OR IGNORE INTO quote_themes (quote_id, theme_id) VALUES (?, ?)",
				q.ID, tid)
		}
		for _, eid := range q.EvidenceIDs {
			tx.Exec("INSERT OR IGNORE INTO quote_evidence (quote_id, evidence_id) VALUES (?, ?)",
				q.ID, eid)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("failed to commit seed data: %v", err)
	}
	log.Println("database seeded with perennial wisdom")
}
