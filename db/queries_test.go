package db_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	"perennial-wisdom/db"
)

// testDB creates an in-memory SQLite database, runs migrations,
// and seeds it with a known data set for deterministic tests.
func testDB(t *testing.T) *sqlx.DB {
	t.Helper()

	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	conn.Exec("PRAGMA foreign_keys=ON")

	database := sqlx.NewDb(conn, "sqlite")

	// Run migrations
	db.Migrate(database)

	// Seed test data directly (not using store seed to keep test-specific)
	seedTestData(t, database)

	return database
}

func seedTestData(t *testing.T, database *sqlx.DB) {
	t.Helper()
	tx, err := database.Beginx()
	if err != nil {
		t.Fatalf("failed to begin tx: %v", err)
	}
	defer tx.Rollback()

	// Philosophies
	stoicPrinciples, _ := json.Marshal([]string{"Virtue", "Acceptance", "Dichotomy of control"})
	buddhistPrinciples, _ := json.Marshal([]string{"Four Noble Truths", "Eightfold Path"})

	tx.Exec("INSERT INTO philosophies (id, name, origin, core_principles) VALUES (?, ?, ?, ?)",
		"stoic", "Stoicism", "Ancient Greece", string(stoicPrinciples))
	tx.Exec("INSERT INTO philosophies (id, name, origin, core_principles) VALUES (?, ?, ?, ?)",
		"buddhist", "Buddhism", "Ancient India", string(buddhistPrinciples))

	// Philosophy relations
	tx.Exec("INSERT INTO philosophy_relations (philosophy_id, related_id) VALUES (?, ?)", "stoic", "buddhist")

	// Philosophers
	epictetusTeachings, _ := json.Marshal([]string{"Dichotomy of control", "Amor fati"})
	buddhaTeachings, _ := json.Marshal([]string{"Middle Way", "Non-attachment"})

	tx.Exec("INSERT INTO philosophers (id, name, philosophy_id, era, bio, key_teachings) VALUES (?, ?, ?, ?, ?, ?)",
		"epictetus", "Epictetus", "stoic", "55–135 CE", "Former slave turned master", string(epictetusTeachings))
	tx.Exec("INSERT INTO philosophers (id, name, philosophy_id, era, bio, key_teachings) VALUES (?, ?, ?, ?, ?, ?)",
		"buddha", "Siddhartha Gautama", "buddhist", "563–483 BCE", "The Awakened One", string(buddhaTeachings))

	// Themes
	tx.Exec("INSERT INTO themes (id, name, description) VALUES (?, ?, ?)",
		"control", "Sphere of Control", "Focus only on what you can influence")
	tx.Exec("INSERT INTO themes (id, name, description) VALUES (?, ?, ?)",
		"impermanence", "Impermanence", "All things change")

	// Theme-philosophy links
	tx.Exec("INSERT INTO theme_philosophies (theme_id, philosophy_id) VALUES (?, ?)", "control", "stoic")
	tx.Exec("INSERT INTO theme_philosophies (theme_id, philosophy_id) VALUES (?, ?)", "control", "buddhist")
	tx.Exec("INSERT INTO theme_philosophies (theme_id, philosophy_id) VALUES (?, ?)", "impermanence", "buddhist")

	// Evidence
	tx.Exec("INSERT INTO evidence (id, title, finding, field, source) VALUES (?, ?, ?, ?, ?)",
		"neuro-control", "PFC & Control", "Prefrontal cortex activation during reappraisal", "neuroscience", "Davidson 2004")
	tx.Exec("INSERT INTO evidence (id, title, finding, field, source) VALUES (?, ?, ?, ?, ?)",
		"neuro-meditation", "Meditation & Brain", "Meditation thickens prefrontal cortex", "neuropsychology", "Lazar 2005")

	// Evidence-theme links
	tx.Exec("INSERT INTO evidence_themes (evidence_id, theme_id) VALUES (?, ?)", "neuro-control", "control")
	tx.Exec("INSERT INTO evidence_themes (evidence_id, theme_id) VALUES (?, ?)", "neuro-meditation", "impermanence")

	// Quotes
	tx.Exec("INSERT INTO quotes (id, text, philosopher_id, philosophy_id, source) VALUES (?, ?, ?, ?, ?)",
		"q1", "It is not things that disturb us, but our judgments about them.", "epictetus", "stoic", "Enchiridion")
	tx.Exec("INSERT INTO quotes (id, text, philosopher_id, philosophy_id, source) VALUES (?, ?, ?, ?, ?)",
		"q2", "All conditioned things are impermanent.", "buddha", "buddhist", "Dhammapada")
	tx.Exec("INSERT INTO quotes (id, text, philosopher_id, philosophy_id, source) VALUES (?, ?, ?, ?, ?)",
		"q3", "We suffer more in imagination than in reality.", "epictetus", "stoic", "Discourses")

	// Quote-theme links
	tx.Exec("INSERT INTO quote_themes (quote_id, theme_id) VALUES (?, ?)", "q1", "control")
	tx.Exec("INSERT INTO quote_themes (quote_id, theme_id) VALUES (?, ?)", "q2", "impermanence")
	tx.Exec("INSERT INTO quote_themes (quote_id, theme_id) VALUES (?, ?)", "q3", "control")

	// Quote-evidence links
	tx.Exec("INSERT INTO quote_evidence (quote_id, evidence_id) VALUES (?, ?)", "q1", "neuro-control")

	if err := tx.Commit(); err != nil {
		t.Fatalf("failed to commit seed data: %v", err)
	}
}

// ---- Quote Queries ----

func TestListQuotes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListQuotes("", "", "")
	if err != nil {
		t.Fatalf("ListQuotes error: %v", err)
	}
	if len(rows) != 3 {
		t.Errorf("expected 3 quotes, got %d", len(rows))
	}
}

func TestListQuotesFilterByPhilosopher(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListQuotes("epictetus", "", "")
	if err != nil {
		t.Fatalf("ListQuotes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 epictetus quotes, got %d", len(rows))
	}
	for _, row := range rows {
		if row.PhilosopherID != "epictetus" {
			t.Errorf("expected philosopher_id=epictetus, got %s", row.PhilosopherID)
		}
	}
}

func TestListQuotesFilterByPhilosophy(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListQuotes("", "buddhist", "")
	if err != nil {
		t.Fatalf("ListQuotes error: %v", err)
	}
	if len(rows) != 1 {
		t.Errorf("expected 1 buddhist quote, got %d", len(rows))
	}
}

func TestListQuotesFilterByTheme(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListQuotes("", "", "control")
	if err != nil {
		t.Fatalf("ListQuotes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 'control' quotes, got %d", len(rows))
	}
}

func TestGetQuote(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.GetQuote("q1")
	if err != nil {
		t.Fatalf("GetQuote error: %v", err)
	}
	if row.PhilosopherName != "Epictetus" {
		t.Errorf("expected philosopher_name=Epictetus, got %s", row.PhilosopherName)
	}
	if row.PhilosophyName != "Stoicism" {
		t.Errorf("expected philosophy_name=Stoicism, got %s", row.PhilosophyName)
	}
}

func TestGetQuoteNotFound(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	_, err := q.GetQuote("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent quote")
	}
}

func TestRandomQuote(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.RandomQuote()
	if err != nil {
		t.Fatalf("RandomQuote error: %v", err)
	}
	if row.ID == "" {
		t.Error("expected a quote ID")
	}
	if row.PhilosopherName == "" {
		t.Error("expected philosopher_name to be populated")
	}
}

func TestQuoteThemes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	themes, err := q.QuoteThemes("q1")
	if err != nil {
		t.Fatalf("QuoteThemes error: %v", err)
	}
	if len(themes) != 1 {
		t.Errorf("expected 1 theme for q1, got %d", len(themes))
	}
	if themes[0].ID != "control" {
		t.Errorf("expected theme=control, got %s", themes[0].ID)
	}
}

func TestQuoteEvidence(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	evidence, err := q.QuoteEvidence("q1")
	if err != nil {
		t.Fatalf("QuoteEvidence error: %v", err)
	}
	if len(evidence) != 1 {
		t.Errorf("expected 1 evidence for q1, got %d", len(evidence))
	}
}

// ---- Philosopher Queries ----

func TestListPhilosophers(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListPhilosophers("")
	if err != nil {
		t.Fatalf("ListPhilosophers error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 philosophers, got %d", len(rows))
	}
}

func TestListPhilosophersFilterByPhilosophy(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListPhilosophers("stoic")
	if err != nil {
		t.Fatalf("ListPhilosophers error: %v", err)
	}
	if len(rows) != 1 {
		t.Errorf("expected 1 stoic philosopher, got %d", len(rows))
	}
	if rows[0].Name != "Epictetus" {
		t.Errorf("expected Epictetus, got %s", rows[0].Name)
	}
}

func TestGetPhilosopher(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.GetPhilosopher("epictetus")
	if err != nil {
		t.Fatalf("GetPhilosopher error: %v", err)
	}
	if row.Name != "Epictetus" {
		t.Errorf("expected name=Epictetus, got %s", row.Name)
	}
	if row.PhilosophyName != "Stoicism" {
		t.Errorf("expected philosophy_name=Stoicism, got %s", row.PhilosophyName)
	}

	// Test Teachings() method
	teachings := row.Teachings()
	if len(teachings) != 2 {
		t.Errorf("expected 2 teachings, got %d", len(teachings))
	}
}

func TestGetPhilosopherNotFound(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	_, err := q.GetPhilosopher("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent philosopher")
	}
}

func TestPhilosopherQuotes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.PhilosopherQuotes("epictetus")
	if err != nil {
		t.Fatalf("PhilosopherQuotes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 quotes for epictetus, got %d", len(rows))
	}
}

// ---- Philosophy Queries ----

func TestListPhilosophies(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListPhilosophies()
	if err != nil {
		t.Fatalf("ListPhilosophies error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 philosophies, got %d", len(rows))
	}
}

func TestGetPhilosophy(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.GetPhilosophy("stoic")
	if err != nil {
		t.Fatalf("GetPhilosophy error: %v", err)
	}
	if row.Name != "Stoicism" {
		t.Errorf("expected name=Stoicism, got %s", row.Name)
	}

	// Test Principles() method
	principles := row.Principles()
	if len(principles) != 3 {
		t.Errorf("expected 3 principles, got %d", len(principles))
	}
}

func TestGetPhilosophyNotFound(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	_, err := q.GetPhilosophy("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent philosophy")
	}
}

func TestPhilosophyRelated(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	names, err := q.PhilosophyRelated("stoic")
	if err != nil {
		t.Fatalf("PhilosophyRelated error: %v", err)
	}
	if len(names) != 1 {
		t.Errorf("expected 1 related philosophy, got %d", len(names))
	}
	if names[0] != "Buddhism" {
		t.Errorf("expected Buddhism, got %s", names[0])
	}
}

func TestPhilosophyPhilosophers(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.PhilosophyPhilosophers("stoic")
	if err != nil {
		t.Fatalf("PhilosophyPhilosophers error: %v", err)
	}
	if len(rows) != 1 {
		t.Errorf("expected 1 stoic philosopher, got %d", len(rows))
	}
}

func TestPhilosophyQuotes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.PhilosophyQuotes("stoic")
	if err != nil {
		t.Fatalf("PhilosophyQuotes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 stoic quotes, got %d", len(rows))
	}
}

// ---- Theme Queries ----

func TestListThemes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListThemes()
	if err != nil {
		t.Fatalf("ListThemes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 themes, got %d", len(rows))
	}
}

func TestGetTheme(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.GetTheme("control")
	if err != nil {
		t.Fatalf("GetTheme error: %v", err)
	}
	if row.Name != "Sphere of Control" {
		t.Errorf("expected name=Sphere of Control, got %s", row.Name)
	}
}

func TestGetThemeNotFound(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	_, err := q.GetTheme("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent theme")
	}
}

func TestThemePhilosophies(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	names, err := q.ThemePhilosophies("control")
	if err != nil {
		t.Fatalf("ThemePhilosophies error: %v", err)
	}
	if len(names) != 2 {
		t.Errorf("expected 2 philosophies for control theme, got %d", len(names))
	}
}

func TestThemeQuotes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ThemeQuotes("control")
	if err != nil {
		t.Fatalf("ThemeQuotes error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 quotes for control theme, got %d", len(rows))
	}
}

func TestThemeEvidence(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ThemeEvidence("control")
	if err != nil {
		t.Fatalf("ThemeEvidence error: %v", err)
	}
	if len(rows) != 1 {
		t.Errorf("expected 1 evidence for control theme, got %d", len(rows))
	}
}

// ---- Evidence Queries ----

func TestListEvidence(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListEvidence("")
	if err != nil {
		t.Fatalf("ListEvidence error: %v", err)
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 evidence, got %d", len(rows))
	}
}

func TestListEvidenceFilterByField(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	rows, err := q.ListEvidence("neuroscience")
	if err != nil {
		t.Fatalf("ListEvidence error: %v", err)
	}
	if len(rows) != 1 {
		t.Errorf("expected 1 neuroscience evidence, got %d", len(rows))
	}
}

func TestGetEvidence(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	row, err := q.GetEvidence("neuro-control")
	if err != nil {
		t.Fatalf("GetEvidence error: %v", err)
	}
	if row.Title != "PFC & Control" {
		t.Errorf("expected title=PFC & Control, got %s", row.Title)
	}
}

func TestGetEvidenceNotFound(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	_, err := q.GetEvidence("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent evidence")
	}
}

func TestEvidenceThemes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	themes, err := q.EvidenceThemes("neuro-control")
	if err != nil {
		t.Fatalf("EvidenceThemes error: %v", err)
	}
	if len(themes) != 1 {
		t.Errorf("expected 1 theme for neuro-control, got %d", len(themes))
	}
}

func TestEvidenceQuotes(t *testing.T) {
	database := testDB(t)
	defer database.Close()
	q := db.NewQueries(database)

	quotes, err := q.EvidenceQuotes("neuro-control")
	if err != nil {
		t.Fatalf("EvidenceQuotes error: %v", err)
	}
	if len(quotes) != 1 {
		t.Errorf("expected 1 quote citing neuro-control, got %d", len(quotes))
	}
}

// ---- Migration Idempotency ----

func TestMigrateIdempotent(t *testing.T) {
	conn, _ := sql.Open("sqlite", ":memory:")
	database := sqlx.NewDb(conn, "sqlite")
	defer database.Close()

	// Run twice — should not panic or error
	db.Migrate(database)
	db.Migrate(database)
}
