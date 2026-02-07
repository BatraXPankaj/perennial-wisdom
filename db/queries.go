package db

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

// Queries provides all database queries for the application.
// Explicit SQL — no ORM, no reflection, fully greppable.
type Queries struct {
	db *sqlx.DB
}

// NewQueries creates a Queries instance with explicit DB dependency.
func NewQueries(db *sqlx.DB) *Queries {
	return &Queries{db: db}
}

// --- Row types (what the DB returns) ---

// QuoteRow is a single quote row from the database.
type QuoteRow struct {
	ID             string `db:"id" json:"id"`
	Text           string `db:"text" json:"text"`
	PhilosopherID  string `db:"philosopher_id" json:"philosopher_id"`
	PhilosophyID   string `db:"philosophy_id" json:"philosophy_id"`
	Source         string `db:"source" json:"source"`
	PhilosopherName string `db:"philosopher_name" json:"philosopher_name,omitempty"`
	PhilosophyName  string `db:"philosophy_name" json:"philosophy_name,omitempty"`
}

// PhilosopherRow is a single philosopher row.
type PhilosopherRow struct {
	ID            string `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	PhilosophyID  string `db:"philosophy_id" json:"philosophy_id"`
	Era           string `db:"era" json:"era"`
	Bio           string `db:"bio" json:"bio"`
	KeyTeachings  string `db:"key_teachings" json:"-"`
	PhilosophyName string `db:"philosophy_name" json:"philosophy_name,omitempty"`
}

// Teachings returns the parsed key_teachings JSON array.
func (p PhilosopherRow) Teachings() []string {
	var t []string
	json.Unmarshal([]byte(p.KeyTeachings), &t)
	return t
}

// PhilosophyRow is a single philosophy row.
type PhilosophyRow struct {
	ID             string `db:"id" json:"id"`
	Name           string `db:"name" json:"name"`
	Origin         string `db:"origin" json:"origin"`
	CorePrinciples string `db:"core_principles" json:"-"`
}

// Principles returns the parsed core_principles JSON array.
func (p PhilosophyRow) Principles() []string {
	var pr []string
	json.Unmarshal([]byte(p.CorePrinciples), &pr)
	return pr
}

// ThemeRow is a single theme row.
type ThemeRow struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

// EvidenceRow is a single evidence row.
type EvidenceRow struct {
	ID      string `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Finding string `db:"finding" json:"finding"`
	Field   string `db:"field" json:"field"`
	Source  string `db:"source" json:"source"`
}

// --- Queries ---

// ListQuotes returns quotes with optional filters.
func (q *Queries) ListQuotes(philosopher, philosophy, theme string) ([]QuoteRow, error) {
	query := `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		WHERE 1=1`
	args := []any{}

	if philosopher != "" {
		query += " AND q.philosopher_id = ?"
		args = append(args, philosopher)
	}
	if philosophy != "" {
		query += " AND q.philosophy_id = ?"
		args = append(args, philosophy)
	}
	if theme != "" {
		query += " AND q.id IN (SELECT quote_id FROM quote_themes WHERE theme_id = ?)"
		args = append(args, theme)
	}

	var rows []QuoteRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetQuote returns a single quote by ID.
func (q *Queries) GetQuote(id string) (QuoteRow, error) {
	var row QuoteRow
	err := q.db.Get(&row, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		WHERE q.id = ?`, id)
	return row, err
}

// RandomQuote returns a random quote.
func (q *Queries) RandomQuote() (QuoteRow, error) {
	var row QuoteRow
	err := q.db.Get(&row, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		ORDER BY RANDOM() LIMIT 1`)
	return row, err
}

// QuoteThemes returns theme names for a quote.
func (q *Queries) QuoteThemes(quoteID string) ([]ThemeRow, error) {
	var rows []ThemeRow
	err := q.db.Select(&rows, `SELECT t.id, t.name, t.description FROM themes t
		JOIN quote_themes qt ON t.id = qt.theme_id
		WHERE qt.quote_id = ?`, quoteID)
	return rows, err
}

// QuoteEvidence returns evidence for a quote.
func (q *Queries) QuoteEvidence(quoteID string) ([]EvidenceRow, error) {
	var rows []EvidenceRow
	err := q.db.Select(&rows, `SELECT e.id, e.title, e.finding, e.field, e.source FROM evidence e
		JOIN quote_evidence qe ON e.id = qe.evidence_id
		WHERE qe.quote_id = ?`, quoteID)
	return rows, err
}

// ListPhilosophers returns philosophers, optionally filtered by philosophy.
func (q *Queries) ListPhilosophers(philosophy string) ([]PhilosopherRow, error) {
	query := `SELECT p.id, p.name, p.philosophy_id, p.era, p.bio, p.key_teachings,
		py.name AS philosophy_name
		FROM philosophers p
		JOIN philosophies py ON p.philosophy_id = py.id
		WHERE 1=1`
	args := []any{}

	if philosophy != "" {
		query += " AND p.philosophy_id = ?"
		args = append(args, philosophy)
	}

	var rows []PhilosopherRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetPhilosopher returns a single philosopher by ID.
func (q *Queries) GetPhilosopher(id string) (PhilosopherRow, error) {
	var row PhilosopherRow
	err := q.db.Get(&row, `SELECT p.id, p.name, p.philosophy_id, p.era, p.bio, p.key_teachings,
		py.name AS philosophy_name
		FROM philosophers p
		JOIN philosophies py ON p.philosophy_id = py.id
		WHERE p.id = ?`, id)
	return row, err
}

// PhilosopherQuotes returns all quotes by a philosopher.
func (q *Queries) PhilosopherQuotes(philosopherID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		WHERE q.philosopher_id = ?`, philosopherID)
	return rows, err
}

// ListPhilosophies returns all philosophy schools.
func (q *Queries) ListPhilosophies() ([]PhilosophyRow, error) {
	var rows []PhilosophyRow
	err := q.db.Select(&rows, "SELECT id, name, origin, core_principles FROM philosophies")
	return rows, err
}

// GetPhilosophy returns a single philosophy by ID.
func (q *Queries) GetPhilosophy(id string) (PhilosophyRow, error) {
	var row PhilosophyRow
	err := q.db.Get(&row, "SELECT id, name, origin, core_principles FROM philosophies WHERE id = ?", id)
	return row, err
}

// PhilosophyRelated returns related philosophy names.
func (q *Queries) PhilosophyRelated(id string) ([]string, error) {
	var names []string
	err := q.db.Select(&names, `SELECT py.name FROM philosophies py
		JOIN philosophy_relations pr ON py.id = pr.related_id
		WHERE pr.philosophy_id = ?`, id)
	return names, err
}

// PhilosophyPhilosophers returns philosophers of a school.
func (q *Queries) PhilosophyPhilosophers(philosophyID string) ([]PhilosopherRow, error) {
	var rows []PhilosopherRow
	err := q.db.Select(&rows, `SELECT p.id, p.name, p.philosophy_id, p.era, p.bio, p.key_teachings,
		py.name AS philosophy_name
		FROM philosophers p
		JOIN philosophies py ON p.philosophy_id = py.id
		WHERE p.philosophy_id = ?`, philosophyID)
	return rows, err
}

// PhilosophyQuotes returns quotes from a philosophy.
func (q *Queries) PhilosophyQuotes(philosophyID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		WHERE q.philosophy_id = ?`, philosophyID)
	return rows, err
}

// ListThemes returns all themes.
func (q *Queries) ListThemes() ([]ThemeRow, error) {
	var rows []ThemeRow
	err := q.db.Select(&rows, "SELECT id, name, description FROM themes")
	return rows, err
}

// GetTheme returns a single theme by ID.
func (q *Queries) GetTheme(id string) (ThemeRow, error) {
	var row ThemeRow
	err := q.db.Get(&row, "SELECT id, name, description FROM themes WHERE id = ?", id)
	return row, err
}

// ThemePhilosophies returns philosophy names for a theme.
func (q *Queries) ThemePhilosophies(themeID string) ([]string, error) {
	var names []string
	err := q.db.Select(&names, `SELECT py.name FROM philosophies py
		JOIN theme_philosophies tp ON py.id = tp.philosophy_id
		WHERE tp.theme_id = ?`, themeID)
	return names, err
}

// ThemeQuotes returns quotes that reference a theme.
func (q *Queries) ThemeQuotes(themeID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		JOIN quote_themes qt ON q.id = qt.quote_id
		WHERE qt.theme_id = ?`, themeID)
	return rows, err
}

// ThemeEvidence returns evidence supporting a theme.
func (q *Queries) ThemeEvidence(themeID string) ([]EvidenceRow, error) {
	var rows []EvidenceRow
	err := q.db.Select(&rows, `SELECT e.id, e.title, e.finding, e.field, e.source FROM evidence e
		JOIN evidence_themes et ON e.id = et.evidence_id
		WHERE et.theme_id = ?`, themeID)
	return rows, err
}

// ListEvidence returns all evidence, optionally filtered by field.
func (q *Queries) ListEvidence(field string) ([]EvidenceRow, error) {
	query := "SELECT id, title, finding, field, source FROM evidence WHERE 1=1"
	args := []any{}

	if field != "" {
		query += " AND field = ?"
		args = append(args, field)
	}

	var rows []EvidenceRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetEvidence returns a single evidence entry by ID.
func (q *Queries) GetEvidence(id string) (EvidenceRow, error) {
	var row EvidenceRow
	err := q.db.Get(&row, "SELECT id, title, finding, field, source FROM evidence WHERE id = ?", id)
	return row, err
}

// EvidenceThemes returns themes linked to evidence.
func (q *Queries) EvidenceThemes(evidenceID string) ([]ThemeRow, error) {
	var rows []ThemeRow
	err := q.db.Select(&rows, `SELECT t.id, t.name, t.description FROM themes t
		JOIN evidence_themes et ON t.id = et.theme_id
		WHERE et.evidence_id = ?`, evidenceID)
	return rows, err
}

// EvidenceQuotes returns quotes that cite evidence.
func (q *Queries) EvidenceQuotes(evidenceID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.text, q.philosopher_id, q.philosophy_id, q.source,
		ph.name AS philosopher_name, py.name AS philosophy_name
		FROM quotes q
		JOIN philosophers ph ON q.philosopher_id = ph.id
		JOIN philosophies py ON q.philosophy_id = py.id
		JOIN quote_evidence qe ON q.id = qe.quote_id
		WHERE qe.evidence_id = ?`, evidenceID)
	return rows, err
}
