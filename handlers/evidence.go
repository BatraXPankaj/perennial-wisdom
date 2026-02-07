package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

// EvidenceHandler serves neuroscience/neuropsychology evidence endpoints.
// Bridges ancient contemplative insight with modern empirical findings.
type EvidenceHandler struct {
	store *store.Store
}

// NewEvidenceHandler creates an EvidenceHandler with explicit store dependency.
func NewEvidenceHandler(s *store.Store) *EvidenceHandler {
	return &EvidenceHandler{store: s}
}

// List returns all scientific evidence, optionally filtered by field:
//   - ?field=neuroscience
//   - ?field=neuropsychology
//   - ?field=psychology
func (h *EvidenceHandler) List(c *gin.Context) {
	field := c.Query("field")

	var results []models.Evidence
	for _, e := range h.store.Evidence {
		if field != "" && e.Field != field {
			continue
		}
		results = append(results, e)
	}

	c.JSON(http.StatusOK, gin.H{"evidence": results, "count": len(results)})
}

// Get returns a single evidence entry by ID, with linked themes and quotes.
func (h *EvidenceHandler) Get(c *gin.Context) {
	id := c.Param("id")
	e, ok := h.store.Evidence[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "evidence not found"})
		return
	}

	// Gather related themes
	var themes []models.Theme
	for _, tid := range e.ThemeIDs {
		if t, ok := h.store.Themes[tid]; ok {
			themes = append(themes, t)
		}
	}

	// Gather quotes that cite this evidence
	var quotes []gin.H
	for _, q := range h.store.Quotes {
		if contains(q.EvidenceIDs, id) {
			quotes = append(quotes, gin.H{
				"quote":       q.Text,
				"philosopher": h.store.Philosophers[q.PhilosopherID].Name,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"evidence": e,
		"themes":   themes,
		"quotes":   quotes,
	})
}
