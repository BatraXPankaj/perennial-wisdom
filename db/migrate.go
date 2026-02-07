package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// schema is the complete database schema.
// Flat, readable SQL — no ORM magic, fully greppable.
const schema = `
CREATE TABLE IF NOT EXISTS philosophies (
	id              TEXT PRIMARY KEY,
	name            TEXT NOT NULL,
	origin          TEXT NOT NULL,
	core_principles TEXT NOT NULL  -- JSON array
);

CREATE TABLE IF NOT EXISTS philosophy_relations (
	philosophy_id TEXT NOT NULL REFERENCES philosophies(id),
	related_id    TEXT NOT NULL REFERENCES philosophies(id),
	PRIMARY KEY (philosophy_id, related_id)
);

CREATE TABLE IF NOT EXISTS philosophers (
	id             TEXT PRIMARY KEY,
	name           TEXT NOT NULL,
	philosophy_id  TEXT NOT NULL REFERENCES philosophies(id),
	era            TEXT NOT NULL,
	bio            TEXT NOT NULL,
	key_teachings  TEXT NOT NULL  -- JSON array
);

CREATE TABLE IF NOT EXISTS themes (
	id          TEXT PRIMARY KEY,
	name        TEXT NOT NULL,
	description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS theme_philosophies (
	theme_id      TEXT NOT NULL REFERENCES themes(id),
	philosophy_id TEXT NOT NULL REFERENCES philosophies(id),
	PRIMARY KEY (theme_id, philosophy_id)
);

CREATE TABLE IF NOT EXISTS evidence (
	id      TEXT PRIMARY KEY,
	title   TEXT NOT NULL,
	finding TEXT NOT NULL,
	field   TEXT NOT NULL,
	source  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS evidence_themes (
	evidence_id TEXT NOT NULL REFERENCES evidence(id),
	theme_id    TEXT NOT NULL REFERENCES themes(id),
	PRIMARY KEY (evidence_id, theme_id)
);

CREATE TABLE IF NOT EXISTS quotes (
	id             TEXT PRIMARY KEY,
	text           TEXT NOT NULL,
	philosopher_id TEXT NOT NULL REFERENCES philosophers(id),
	philosophy_id  TEXT NOT NULL REFERENCES philosophies(id),
	source         TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS quote_themes (
	quote_id TEXT NOT NULL REFERENCES quotes(id),
	theme_id TEXT NOT NULL REFERENCES themes(id),
	PRIMARY KEY (quote_id, theme_id)
);

CREATE TABLE IF NOT EXISTS quote_evidence (
	quote_id    TEXT NOT NULL REFERENCES quotes(id),
	evidence_id TEXT NOT NULL REFERENCES evidence(id),
	PRIMARY KEY (quote_id, evidence_id)
);
`

// Migrate runs the schema creation.
// Idempotent — safe to run on every startup.
func Migrate(db *sqlx.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("failed to run migration: %v", err)
	}
	log.Println("database schema migrated successfully")
}
