package db

import (
	"database/sql"
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
	ID                    string         `db:"id" json:"id"`
	Title                 sql.NullString `db:"title" json:"title,omitempty"`
	Slug                  sql.NullString `db:"slug" json:"slug,omitempty"`
	Text                  string         `db:"text" json:"text"`
	TextScholarly         sql.NullString `db:"text_scholarly" json:"text_scholarly,omitempty"`
	PhilosopherID         sql.NullString `db:"philosopher_id" json:"philosopher_id"`
	TraditionID           sql.NullString `db:"tradition_id" json:"tradition_id"`
	SourceWork            sql.NullString `db:"source_work" json:"source_work,omitempty"`
	SourceLocation        sql.NullString `db:"source_location" json:"source_location,omitempty"`
	OriginalScript        sql.NullString `db:"original_script" json:"original_script,omitempty"`
	ExpositionBrief       sql.NullString `db:"exposition_brief" json:"exposition_brief,omitempty"`
	ExpositionStandard    sql.NullString `db:"exposition_standard" json:"exposition_standard,omitempty"`
	ExpositionScholarly   sql.NullString `db:"exposition_scholarly" json:"exposition_scholarly,omitempty"`
	ReflectionPrompt      sql.NullString `db:"reflection_prompt" json:"reflection_prompt,omitempty"`
	ModernReinterpretation sql.NullString `db:"modern_reinterpretation" json:"modern_reinterpretation,omitempty"`
	Meta                  []byte         `db:"meta" json:"-"`
	PhilosopherName       sql.NullString `db:"philosopher_name" json:"philosopher_name,omitempty"`
	TraditionName         sql.NullString `db:"tradition_name" json:"tradition_name,omitempty"`
}

// GetTitle returns the title or a default.
func (q QuoteRow) GetTitle() string {
	if q.Title.Valid {
		return q.Title.String
	}
	return ""
}

// GetMeta returns parsed meta JSONB.
func (q QuoteRow) GetMeta() map[string]any {
	if len(q.Meta) == 0 {
		return nil
	}
	var m map[string]any
	json.Unmarshal(q.Meta, &m)
	return m
}

// PhilosopherRow is a single philosopher row.
type PhilosopherRow struct {
	ID            string         `db:"id" json:"id"`
	Name          string         `db:"name" json:"name"`
	TraditionID   sql.NullString `db:"tradition_id" json:"tradition_id"`
	Era           sql.NullString `db:"era" json:"era"`
	Bio           sql.NullString `db:"bio" json:"bio"`
	KeyTeachings  []byte         `db:"key_teachings" json:"-"`
	TraditionName sql.NullString `db:"tradition_name" json:"tradition_name,omitempty"`
}

// Teachings returns the parsed key_teachings JSONB array.
func (p PhilosopherRow) Teachings() []string {
	var t []string
	json.Unmarshal(p.KeyTeachings, &t)
	return t
}

// TraditionRow is a single tradition (philosophy school) row.
type TraditionRow struct {
	ID             string         `db:"id" json:"id"`
	Name           string         `db:"name" json:"name"`
	Origin         sql.NullString `db:"origin" json:"origin"`
	CorePrinciples []byte         `db:"core_principles" json:"-"`
}

// Principles returns the parsed core_principles JSONB array.
func (t TraditionRow) Principles() []string {
	var pr []string
	json.Unmarshal(t.CorePrinciples, &pr)
	return pr
}

// ThemeRow is a single theme row.
type ThemeRow struct {
	ID          string         `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
}

// EvidenceRow is a single evidence row.
type EvidenceRow struct {
	ID               string         `db:"id" json:"id"`
	Title            string         `db:"title" json:"title"`
	Finding          sql.NullString `db:"finding" json:"finding"`
	Field            sql.NullString `db:"field" json:"field"`
	Citation         sql.NullString `db:"citation" json:"citation"`
	EvidenceStrength sql.NullString `db:"evidence_strength" json:"evidence_strength"`
}

// --- Queries ---

// ListQuotes returns quotes with optional filters.
func (q *Queries) ListQuotes(philosopher, tradition, theme string) ([]QuoteRow, error) {
	query := `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		WHERE 1=1`
	args := []any{}
	argNum := 1

	if philosopher != "" {
		query += " AND q.philosopher_id = $" + itoa(argNum)
		args = append(args, philosopher)
		argNum++
	}
	if tradition != "" {
		query += " AND q.tradition_id = $" + itoa(argNum)
		args = append(args, tradition)
		argNum++
	}
	if theme != "" {
		query += " AND q.id IN (SELECT quote_id FROM quote_themes WHERE theme_id = $" + itoa(argNum) + ")"
		args = append(args, theme)
		argNum++
	}

	query += " ORDER BY q.created_at DESC"

	var rows []QuoteRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetQuote returns a single quote by ID or slug.
func (q *Queries) GetQuote(idOrSlug string) (QuoteRow, error) {
	var row QuoteRow
	err := q.db.Get(&row, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.original_script, q.exposition_brief, q.exposition_standard, q.exposition_scholarly,
		q.reflection_prompt, q.modern_reinterpretation, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		WHERE q.id = $1 OR q.slug = $1`, idOrSlug)
	return row, err
}

// RandomQuote returns a random quote.
func (q *Queries) RandomQuote() (QuoteRow, error) {
	var row QuoteRow
	err := q.db.Get(&row, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		ORDER BY RANDOM() LIMIT 1`)
	return row, err
}

// QuoteThemes returns theme names for a quote.
func (q *Queries) QuoteThemes(quoteID string) ([]ThemeRow, error) {
	var rows []ThemeRow
	err := q.db.Select(&rows, `SELECT t.id, t.name, t.description FROM themes t
		JOIN quote_themes qt ON t.id = qt.theme_id
		WHERE qt.quote_id = $1`, quoteID)
	return rows, err
}

// QuoteEvidence returns evidence for a quote.
func (q *Queries) QuoteEvidence(quoteID string) ([]EvidenceRow, error) {
	var rows []EvidenceRow
	err := q.db.Select(&rows, `SELECT e.id, e.title, e.finding, e.field, e.citation, e.evidence_strength 
		FROM evidence e
		JOIN quote_evidence qe ON e.id = qe.evidence_id
		WHERE qe.quote_id = $1`, quoteID)
	return rows, err
}

// ListPhilosophers returns philosophers, optionally filtered by tradition.
func (q *Queries) ListPhilosophers(tradition string) ([]PhilosopherRow, error) {
	query := `SELECT p.id, p.name, p.tradition_id, p.era, p.bio, p.key_teachings,
		t.name AS tradition_name
		FROM philosophers p
		LEFT JOIN traditions t ON p.tradition_id = t.id
		WHERE 1=1`
	args := []any{}

	if tradition != "" {
		query += " AND p.tradition_id = $1"
		args = append(args, tradition)
	}

	var rows []PhilosopherRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetPhilosopher returns a single philosopher by ID.
func (q *Queries) GetPhilosopher(id string) (PhilosopherRow, error) {
	var row PhilosopherRow
	err := q.db.Get(&row, `SELECT p.id, p.name, p.tradition_id, p.era, p.bio, p.key_teachings,
		t.name AS tradition_name
		FROM philosophers p
		LEFT JOIN traditions t ON p.tradition_id = t.id
		WHERE p.id = $1`, id)
	return row, err
}

// PhilosopherQuotes returns all quotes by a philosopher.
func (q *Queries) PhilosopherQuotes(philosopherID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		WHERE q.philosopher_id = $1`, philosopherID)
	return rows, err
}

// ListTraditions returns all tradition schools.
func (q *Queries) ListTraditions() ([]TraditionRow, error) {
	var rows []TraditionRow
	err := q.db.Select(&rows, "SELECT id, name, origin, core_principles FROM traditions ORDER BY name")
	return rows, err
}

// GetTradition returns a single tradition by ID.
func (q *Queries) GetTradition(id string) (TraditionRow, error) {
	var row TraditionRow
	err := q.db.Get(&row, "SELECT id, name, origin, core_principles FROM traditions WHERE id = $1", id)
	return row, err
}

// TraditionPhilosophers returns philosophers of a school.
func (q *Queries) TraditionPhilosophers(traditionID string) ([]PhilosopherRow, error) {
	var rows []PhilosopherRow
	err := q.db.Select(&rows, `SELECT p.id, p.name, p.tradition_id, p.era, p.bio, p.key_teachings,
		t.name AS tradition_name
		FROM philosophers p
		LEFT JOIN traditions t ON p.tradition_id = t.id
		WHERE p.tradition_id = $1`, traditionID)
	return rows, err
}

// TraditionQuotes returns quotes from a tradition.
func (q *Queries) TraditionQuotes(traditionID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		WHERE q.tradition_id = $1`, traditionID)
	return rows, err
}

// ListThemes returns all themes.
func (q *Queries) ListThemes() ([]ThemeRow, error) {
	var rows []ThemeRow
	err := q.db.Select(&rows, "SELECT id, name, description FROM themes ORDER BY name")
	return rows, err
}

// GetTheme returns a single theme by ID.
func (q *Queries) GetTheme(id string) (ThemeRow, error) {
	var row ThemeRow
	err := q.db.Get(&row, "SELECT id, name, description FROM themes WHERE id = $1", id)
	return row, err
}

// ThemeQuotes returns quotes that reference a theme.
func (q *Queries) ThemeQuotes(themeID string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		JOIN quote_themes qt ON q.id = qt.quote_id
		WHERE qt.theme_id = $1`, themeID)
	return rows, err
}

// ListEvidence returns all evidence, optionally filtered by field.
func (q *Queries) ListEvidence(field string) ([]EvidenceRow, error) {
	query := "SELECT id, title, finding, field, citation, evidence_strength FROM evidence WHERE 1=1"
	args := []any{}

	if field != "" {
		query += " AND field = $1"
		args = append(args, field)
	}

	var rows []EvidenceRow
	err := q.db.Select(&rows, query, args...)
	return rows, err
}

// GetEvidence returns a single evidence entry by ID.
func (q *Queries) GetEvidence(id string) (EvidenceRow, error) {
	var row EvidenceRow
	err := q.db.Get(&row, "SELECT id, title, finding, field, citation, evidence_strength FROM evidence WHERE id = $1", id)
	return row, err
}

// SearchQuotes performs full-text search on quotes.
func (q *Queries) SearchQuotes(query string) ([]QuoteRow, error) {
	var rows []QuoteRow
	err := q.db.Select(&rows, `SELECT q.id, q.title, q.slug, q.text, q.text_scholarly,
		q.philosopher_id, q.tradition_id, q.source_work, q.source_location,
		q.exposition_brief, q.meta,
		ph.name AS philosopher_name, t.name AS tradition_name
		FROM quotes q
		LEFT JOIN philosophers ph ON q.philosopher_id = ph.id
		LEFT JOIN traditions t ON q.tradition_id = t.id
		WHERE to_tsvector('english', COALESCE(q.text, '') || ' ' || COALESCE(q.title, '') || ' ' || COALESCE(q.exposition_brief, ''))
		@@ plainto_tsquery('english', $1)
		LIMIT 50`, query)
	return rows, err
}

// Helper to convert int to string for query building
func itoa(n int) string {
	return string(rune('0' + n))
}
