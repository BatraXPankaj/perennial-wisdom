package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

// PhilosophyHandler serves philosophy/school-related endpoints.
type PhilosophyHandler struct {
	store *store.Store
}

// NewPhilosophyHandler creates a PhilosophyHandler with explicit store dependency.
func NewPhilosophyHandler(s *store.Store) *PhilosophyHandler {
	return &PhilosophyHandler{store: s}
}

// List returns all philosophies/schools.
func (h *PhilosophyHandler) List(c *gin.Context) {
	var results []models.Philosophy
	for _, p := range h.store.Philosophies {
		results = append(results, p)
	}

	c.JSON(http.StatusOK, gin.H{"philosophies": results, "count": len(results)})
}

// Get returns a single philosophy by ID, with its philosophers,
// related philosophies, and a sample of quotes.
func (h *PhilosophyHandler) Get(c *gin.Context) {
	id := c.Param("id")
	p, ok := h.store.Philosophies[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "philosophy not found"})
		return
	}

	// Gather related philosophy names
	var related []string
	for _, rid := range p.RelatedIDs {
		if rp, ok := h.store.Philosophies[rid]; ok {
			related = append(related, rp.Name)
		}
	}

	// Gather philosophers of this school
	var philosophers []models.Philosopher
	for _, ph := range h.store.Philosophers {
		if ph.PhilosophyID == id {
			philosophers = append(philosophers, ph)
		}
	}

	// Gather quotes from this school
	var quotes []models.Quote
	for _, q := range h.store.Quotes {
		if q.PhilosophyID == id {
			quotes = append(quotes, q)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"philosophy":   p,
		"related":      related,
		"philosophers": philosophers,
		"quotes":       quotes,
	})
}
