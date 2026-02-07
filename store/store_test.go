package store_test

import (
	"testing"

	"perennial-wisdom/store"
)

func TestNewStoreNotNil(t *testing.T) {
	s := store.New()
	if s == nil {
		t.Fatal("expected non-nil store")
	}
}

func TestStoreHasQuotes(t *testing.T) {
	s := store.New()
	if len(s.Quotes) == 0 {
		t.Error("expected seed quotes, got 0")
	}
}

func TestStoreHasPhilosophers(t *testing.T) {
	s := store.New()
	if len(s.Philosophers) == 0 {
		t.Error("expected seed philosophers, got 0")
	}
}

func TestStoreHasPhilosophies(t *testing.T) {
	s := store.New()
	if len(s.Philosophies) == 0 {
		t.Error("expected seed philosophies, got 0")
	}
}

func TestStoreHasThemes(t *testing.T) {
	s := store.New()
	if len(s.Themes) == 0 {
		t.Error("expected seed themes, got 0")
	}
}

func TestStoreHasEvidence(t *testing.T) {
	s := store.New()
	if len(s.Evidence) == 0 {
		t.Error("expected seed evidence, got 0")
	}
}

func TestStoreQuoteReferences(t *testing.T) {
	s := store.New()

	for id, q := range s.Quotes {
		// Every quote should reference a valid philosopher
		if _, ok := s.Philosophers[q.PhilosopherID]; !ok {
			t.Errorf("quote %s references unknown philosopher %s", id, q.PhilosopherID)
		}
		// Every quote should reference a valid philosophy
		if _, ok := s.Philosophies[q.PhilosophyID]; !ok {
			t.Errorf("quote %s references unknown philosophy %s", id, q.PhilosophyID)
		}
		// Every theme ID should be valid
		for _, tid := range q.ThemeIDs {
			if _, ok := s.Themes[tid]; !ok {
				t.Errorf("quote %s references unknown theme %s", id, tid)
			}
		}
		// Every evidence ID should be valid
		for _, eid := range q.EvidenceIDs {
			if _, ok := s.Evidence[eid]; !ok {
				t.Errorf("quote %s references unknown evidence %s", id, eid)
			}
		}
	}
}

func TestStorePhilosopherReferences(t *testing.T) {
	s := store.New()

	for id, p := range s.Philosophers {
		if _, ok := s.Philosophies[p.PhilosophyID]; !ok {
			t.Errorf("philosopher %s references unknown philosophy %s", id, p.PhilosophyID)
		}
	}
}

func TestStoreThemeReferences(t *testing.T) {
	s := store.New()

	for id, th := range s.Themes {
		for _, pid := range th.PhilosophyIDs {
			if _, ok := s.Philosophies[pid]; !ok {
				t.Errorf("theme %s references unknown philosophy %s", id, pid)
			}
		}
	}
}

func TestStoreEvidenceReferences(t *testing.T) {
	s := store.New()

	for id, e := range s.Evidence {
		for _, tid := range e.ThemeIDs {
			if _, ok := s.Themes[tid]; !ok {
				t.Errorf("evidence %s references unknown theme %s", id, tid)
			}
		}
	}
}

func TestStorePhilosophyRelatedReferences(t *testing.T) {
	s := store.New()

	for id, p := range s.Philosophies {
		for _, rid := range p.RelatedIDs {
			if _, ok := s.Philosophies[rid]; !ok {
				t.Errorf("philosophy %s references unknown related philosophy %s", id, rid)
			}
		}
	}
}
