package handlers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/db"
)

// Pages serves HTML pages using Go templates + HTMX.
// All dependencies are explicit — queries for data, templates for rendering.
type Pages struct {
	q    *db.Queries
	tmpl *template.Template
}

// NewPages creates a Pages handler with explicit dependencies.
func NewPages(q *db.Queries, tmpl *template.Template) *Pages {
	return &Pages{q: q, tmpl: tmpl}
}

// render is a helper that executes a named template with base layout.
func (p *Pages) render(c *gin.Context, status int, data gin.H) {
	data["Title"] = data["Title"] // ensure Title key exists
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := p.tmpl.ExecuteTemplate(c.Writer, "base", data); err != nil {
		c.String(http.StatusInternalServerError, "template error: %v", err)
	}
}

// Home renders the landing page with a random quote and philosophy grid.
func (p *Pages) Home(c *gin.Context) {
	philosophies, _ := p.q.ListPhilosophies()
	p.render(c, http.StatusOK, gin.H{
		"Title":        "Home",
		"Philosophies": philosophies,
	})
}

// RandomQuotePartial returns just the random quote HTML fragment (for HTMX).
func (p *Pages) RandomQuotePartial(c *gin.Context) {
	quote, err := p.q.RandomQuote()
	if err != nil {
		c.String(http.StatusInternalServerError, "no quotes available")
		return
	}
	c.Header("Content-Type", "text/html; charset=utf-8")
	p.tmpl.ExecuteTemplate(c.Writer, "random-quote", quote)
}

// Quotes renders the quotes listing with filters.
func (p *Pages) Quotes(c *gin.Context) {
	philosophy := c.Query("philosophy")
	theme := c.Query("theme")

	quotes, _ := p.q.ListQuotes("", philosophy, theme)
	philosophies, _ := p.q.ListPhilosophies()
	themes, _ := p.q.ListThemes()

	p.render(c, http.StatusOK, gin.H{
		"Title":        "Quotes",
		"Quotes":       quotes,
		"Philosophies": philosophies,
		"Themes":       themes,
		"Filter": gin.H{
			"Philosophy": philosophy,
			"Theme":      theme,
		},
	})
}

// Philosophers renders the philosophers listing.
func (p *Pages) Philosophers(c *gin.Context) {
	philosophers, _ := p.q.ListPhilosophers("")
	p.render(c, http.StatusOK, gin.H{
		"Title":        "Philosophers",
		"Philosophers": philosophers,
	})
}

// PhilosopherDetail renders a single philosopher page.
func (p *Pages) PhilosopherDetail(c *gin.Context) {
	id := c.Param("id")
	philosopher, err := p.q.GetPhilosopher(id)
	if err != nil {
		c.String(http.StatusNotFound, "philosopher not found")
		return
	}
	quotes, _ := p.q.PhilosopherQuotes(id)

	p.render(c, http.StatusOK, gin.H{
		"Title":       philosopher.Name,
		"Philosopher": philosopher,
		"Teachings":   philosopher.Teachings(),
		"Quotes":      quotes,
	})
}

// Philosophies renders the philosophies listing.
func (p *Pages) Philosophies(c *gin.Context) {
	philosophies, _ := p.q.ListPhilosophies()
	p.render(c, http.StatusOK, gin.H{
		"Title":        "Schools of Wisdom",
		"Philosophies": philosophies,
	})
}

// PhilosophyDetail renders a single philosophy page.
func (p *Pages) PhilosophyDetail(c *gin.Context) {
	id := c.Param("id")
	philosophy, err := p.q.GetPhilosophy(id)
	if err != nil {
		c.String(http.StatusNotFound, "philosophy not found")
		return
	}
	related, _ := p.q.PhilosophyRelated(id)
	philosophers, _ := p.q.PhilosophyPhilosophers(id)
	quotes, _ := p.q.PhilosophyQuotes(id)

	p.render(c, http.StatusOK, gin.H{
		"Title":        philosophy.Name,
		"Philosophy":   philosophy,
		"Principles":   philosophy.Principles(),
		"Related":      related,
		"Philosophers": philosophers,
		"Quotes":       quotes,
	})
}

// Themes renders the themes listing.
func (p *Pages) Themes(c *gin.Context) {
	themes, _ := p.q.ListThemes()
	p.render(c, http.StatusOK, gin.H{
		"Title":  "Perennial Themes",
		"Themes": themes,
	})
}

// ThemeDetail renders a single theme page — the cross-correlation view.
func (p *Pages) ThemeDetail(c *gin.Context) {
	id := c.Param("id")
	theme, err := p.q.GetTheme(id)
	if err != nil {
		c.String(http.StatusNotFound, "theme not found")
		return
	}
	philosophies, _ := p.q.ThemePhilosophies(id)
	quotes, _ := p.q.ThemeQuotes(id)
	evidence, _ := p.q.ThemeEvidence(id)

	p.render(c, http.StatusOK, gin.H{
		"Title":        theme.Name,
		"Theme":        theme,
		"Philosophies": philosophies,
		"Quotes":       quotes,
		"Evidence":     evidence,
	})
}

// Evidence renders the evidence listing.
func (p *Pages) Evidence(c *gin.Context) {
	field := c.Query("field")
	evidence, _ := p.q.ListEvidence(field)
	p.render(c, http.StatusOK, gin.H{
		"Title":    "Scientific Evidence",
		"Evidence": evidence,
		"Filter":   field,
	})
}

// EvidenceDetail renders a single evidence page.
func (p *Pages) EvidenceDetail(c *gin.Context) {
	id := c.Param("id")
	evidence, err := p.q.GetEvidence(id)
	if err != nil {
		c.String(http.StatusNotFound, "evidence not found")
		return
	}
	themes, _ := p.q.EvidenceThemes(id)
	quotes, _ := p.q.EvidenceQuotes(id)

	p.render(c, http.StatusOK, gin.H{
		"Title":    evidence.Title,
		"Evidence": evidence,
		"Themes":   themes,
		"Quotes":   quotes,
	})
}
