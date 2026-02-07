package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

// QuoteHandler serves quote-related endpoints.
// Quotes are the center stage of the perennial wisdom API.
type QuoteHandler struct {
	store *store.Store
}

// NewQuoteHandler creates a QuoteHandler with explicit store dependency.
func NewQuoteHandler(s *store.Store) *QuoteHandler {
	return &QuoteHandler{store: s}
}

// List returns all quotes, with optional filters:
//   - ?philosopher=epictetus
//   - ?philosophy=stoic
//   - ?theme=control
func (h *QuoteHandler) List(c *gin.Context) {
	philosopher := c.Query("philosopher")
	philosophy := c.Query("philosophy")
	theme := c.Query("theme")

	var results []models.Quote
	for _, q := range h.store.Quotes {
		if philosopher != "" && q.PhilosopherID != philosopher {
			continue
		}
		if philosophy != "" && q.PhilosophyID != philosophy {
			continue
		}
		if theme != "" && !contains(q.ThemeIDs, theme) {
			continue
		}
		results = append(results, q)
	}

	c.JSON(http.StatusOK, gin.H{"quotes": results, "count": len(results)})
}

// Get returns a single quote by ID, enriched with philosopher and theme names.
func (h *QuoteHandler) Get(c *gin.Context) {
	id := c.Param("id")
	q, ok := h.store.Quotes[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "quote not found"})
		return
	}

	philosopher := h.store.Philosophers[q.PhilosopherID]

	c.JSON(http.StatusOK, gin.H{
		"quote":       q,
		"philosopher": philosopher.Name,
		"philosophy":  h.store.Philosophies[q.PhilosophyID].Name,
	})
}

// Random returns a random quote. Uses map iteration order
// which is randomized in Go by design.
func (h *QuoteHandler) Random(c *gin.Context) {
	for _, q := range h.store.Quotes {
		philosopher := h.store.Philosophers[q.PhilosopherID]
		c.JSON(http.StatusOK, gin.H{
			"quote":       q,
			"philosopher": philosopher.Name,
			"philosophy":  h.store.Philosophies[q.PhilosophyID].Name,
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "no quotes available"})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
