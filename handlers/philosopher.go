package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

// PhilosopherHandler serves philosopher-related endpoints.
type PhilosopherHandler struct {
	store *store.Store
}

// NewPhilosopherHandler creates a PhilosopherHandler with explicit store dependency.
func NewPhilosopherHandler(s *store.Store) *PhilosopherHandler {
	return &PhilosopherHandler{store: s}
}

// List returns all philosophers, optionally filtered by philosophy:
//   - ?philosophy=stoic
func (h *PhilosopherHandler) List(c *gin.Context) {
	philosophy := c.Query("philosophy")

	var results []models.Philosopher
	for _, p := range h.store.Philosophers {
		if philosophy != "" && p.PhilosophyID != philosophy {
			continue
		}
		results = append(results, p)
	}

	c.JSON(http.StatusOK, gin.H{"philosophers": results, "count": len(results)})
}

// Get returns a single philosopher by ID, with their quotes.
func (h *PhilosopherHandler) Get(c *gin.Context) {
	id := c.Param("id")
	p, ok := h.store.Philosophers[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "philosopher not found"})
		return
	}

	var quotes []models.Quote
	for _, q := range h.store.Quotes {
		if q.PhilosopherID == id {
			quotes = append(quotes, q)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"philosopher": p,
		"philosophy":  h.store.Philosophies[p.PhilosophyID].Name,
		"quotes":      quotes,
	})
}
