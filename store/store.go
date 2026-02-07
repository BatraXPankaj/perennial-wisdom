package store

import "perennial-wisdom/models"

// Store holds all in-memory data, indexed by ID for fast lookup.
// Flat maps — no ORM, no magic, fully transparent to agents and humans.
type Store struct {
	Quotes       map[string]models.Quote
	Philosophers map[string]models.Philosopher
	Philosophies map[string]models.Philosophy
	Themes       map[string]models.Theme
	Evidence     map[string]models.Evidence
}

// New creates a Store pre-loaded with seed data.
// All data is in memory; swap this for a database layer later
// by implementing the same method signatures on a DB-backed store.
func New() *Store {
	s := &Store{
		Quotes:       make(map[string]models.Quote),
		Philosophers: make(map[string]models.Philosopher),
		Philosophies: make(map[string]models.Philosophy),
		Themes:       make(map[string]models.Theme),
		Evidence:     make(map[string]models.Evidence),
	}

	for _, p := range SeedPhilosophies() {
		s.Philosophies[p.ID] = p
	}
	for _, p := range SeedPhilosophers() {
		s.Philosophers[p.ID] = p
	}
	for _, t := range SeedThemes() {
		s.Themes[t.ID] = t
	}
	for _, e := range SeedEvidence() {
		s.Evidence[e.ID] = e
	}
	for _, q := range SeedQuotes() {
		s.Quotes[q.ID] = q
	}

	return s
}
