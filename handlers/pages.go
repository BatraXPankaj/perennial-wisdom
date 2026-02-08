package handlers

import (
	"bytes"
	"html/template"
	"log"
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

// render executes the "base" template with page-specific content.
// Renders to a buffer first to avoid partial HTML on error.
func (p *Pages) render(c *gin.Context, status int, data gin.H) {
	var buf bytes.Buffer
	if err := p.tmpl.ExecuteTemplate(&buf, "base", data); err != nil {
		log.Printf("template error: %v", err)
		c.String(http.StatusInternalServerError, "template error: %v", err)
		return
	}
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(c.Writer)
}

// Home renders the landing page with a random quote and philosophy grid.
func (p *Pages) Home(c *gin.Context) {
	philosophies, err := p.q.ListPhilosophies()
	if err != nil {
		log.Printf("Home: ListPhilosophies error: %v", err)
	}
	p.render(c, http.StatusOK, gin.H{
		"Page":         "home",
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

	quotes, err := p.q.ListQuotes("", philosophy, theme)
	if err != nil {
		log.Printf("Quotes: ListQuotes error: %v", err)
	}
	philosophies, err := p.q.ListPhilosophies()
	if err != nil {
		log.Printf("Quotes: ListPhilosophies error: %v", err)
	}
	themes, err := p.q.ListThemes()
	if err != nil {
		log.Printf("Quotes: ListThemes error: %v", err)
	}

	p.render(c, http.StatusOK, gin.H{
		"Page":         "quotes",
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
	philosophers, err := p.q.ListPhilosophers("")
	if err != nil {
		log.Printf("Philosophers: ListPhilosophers error: %v", err)
	}
	p.render(c, http.StatusOK, gin.H{
		"Page":         "philosophers",
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
	quotes, err := p.q.PhilosopherQuotes(id)
	if err != nil {
		log.Printf("PhilosopherDetail: PhilosopherQuotes error: %v", err)
	}

	p.render(c, http.StatusOK, gin.H{
		"Page":        "philosopher-detail",
		"Title":       philosopher.Name,
		"Philosopher": philosopher,
		"Teachings":   philosopher.Teachings(),
		"Quotes":      quotes,
	})
}

// Philosophies renders the philosophies listing.
func (p *Pages) Philosophies(c *gin.Context) {
	philosophies, err := p.q.ListPhilosophies()
	if err != nil {
		log.Printf("Philosophies: ListPhilosophies error: %v", err)
	}
	p.render(c, http.StatusOK, gin.H{
		"Page":         "philosophies",
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
	related, err := p.q.PhilosophyRelated(id)
	if err != nil {
		log.Printf("PhilosophyDetail: PhilosophyRelated error: %v", err)
	}
	philosophers, err := p.q.PhilosophyPhilosophers(id)
	if err != nil {
		log.Printf("PhilosophyDetail: PhilosophyPhilosophers error: %v", err)
	}
	quotes, err := p.q.PhilosophyQuotes(id)
	if err != nil {
		log.Printf("PhilosophyDetail: PhilosophyQuotes error: %v", err)
	}

	p.render(c, http.StatusOK, gin.H{
		"Page":         "philosophy-detail",
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
	themes, err := p.q.ListThemes()
	if err != nil {
		log.Printf("Themes: ListThemes error: %v", err)
	}
	p.render(c, http.StatusOK, gin.H{
		"Page":   "themes",
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
	philosophies, err := p.q.ThemePhilosophies(id)
	if err != nil {
		log.Printf("ThemeDetail: ThemePhilosophies error: %v", err)
	}
	quotes, err := p.q.ThemeQuotes(id)
	if err != nil {
		log.Printf("ThemeDetail: ThemeQuotes error: %v", err)
	}
	evidence, err := p.q.ThemeEvidence(id)
	if err != nil {
		log.Printf("ThemeDetail: ThemeEvidence error: %v", err)
	}

	p.render(c, http.StatusOK, gin.H{
		"Page":         "theme-detail",
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
	evidence, err := p.q.ListEvidence(field)
	if err != nil {
		log.Printf("Evidence: ListEvidence error: %v", err)
	}
	p.render(c, http.StatusOK, gin.H{
		"Page":     "evidence",
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
	themes, err := p.q.EvidenceThemes(id)
	if err != nil {
		log.Printf("EvidenceDetail: EvidenceThemes error: %v", err)
	}
	quotes, err := p.q.EvidenceQuotes(id)
	if err != nil {
		log.Printf("EvidenceDetail: EvidenceQuotes error: %v", err)
	}

	p.render(c, http.StatusOK, gin.H{
		"Page":     "evidence-detail",
		"Title":    evidence.Title,
		"Evidence": evidence,
		"Themes":   themes,
		"Quotes":   quotes,
	})
}
