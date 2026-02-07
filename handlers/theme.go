package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

// ThemeHandler serves theme-related endpoints.
// Themes are the cross-tradition connective tissue — the "perennial" threads.
type ThemeHandler struct {
	store *store.Store
}

// NewThemeHandler creates a ThemeHandler with explicit store dependency.
func NewThemeHandler(s *store.Store) *ThemeHandler {
	return &ThemeHandler{store: s}
}

// List returns all themes.
func (h *ThemeHandler) List(c *gin.Context) {
	var results []models.Theme
	for _, t := range h.store.Themes {
		results = append(results, t)
	}

	c.JSON(http.StatusOK, gin.H{"themes": results, "count": len(results)})
}

// Get returns a single theme by ID, with quotes across traditions
// that address this theme — the cross-correlation view.
func (h *ThemeHandler) Get(c *gin.Context) {
	id := c.Param("id")
	t, ok := h.store.Themes[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "theme not found"})
		return
	}

	// Gather quotes that reference this theme
	var quotes []gin.H
	for _, q := range h.store.Quotes {
		if contains(q.ThemeIDs, id) {
			quotes = append(quotes, gin.H{
				"quote":       q.Text,
				"philosopher": h.store.Philosophers[q.PhilosopherID].Name,
				"philosophy":  h.store.Philosophies[q.PhilosophyID].Name,
				"source":      q.Source,
			})
		}
	}

	// Gather evidence supporting this theme
	var evidence []models.Evidence
	for _, e := range h.store.Evidence {
		if contains(e.ThemeIDs, id) {
			evidence = append(evidence, e)
		}
	}

	// Gather philosophy names that address this theme
	var philosophies []string
	for _, pid := range t.PhilosophyIDs {
		if p, ok := h.store.Philosophies[pid]; ok {
			philosophies = append(philosophies, p.Name)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"theme":        t,
		"philosophies": philosophies,
		"quotes":       quotes,
		"evidence":     evidence,
	})
}
