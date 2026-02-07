package router

import (
	"html/template"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/db"
	"perennial-wisdom/handlers"
	"perennial-wisdom/store"
)

// Setup creates a Gin engine with all routes wired.
// All dependencies are explicit — no init(), no reflection, no magic.
func Setup(s *store.Store, q *db.Queries, tmpl *template.Template) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", handlers.Health)

	// --- JSON API (existing) ---

	// Quotes — center stage
	qh := handlers.NewQuoteHandler(s)
	r.GET("/api/quotes", qh.List)
	r.GET("/api/quotes/random", qh.Random)
	r.GET("/api/quotes/:id", qh.Get)

	// Philosophers — the teachers
	ph := handlers.NewPhilosopherHandler(s)
	r.GET("/api/philosophers", ph.List)
	r.GET("/api/philosophers/:id", ph.Get)

	// Philosophies — the schools
	pyh := handlers.NewPhilosophyHandler(s)
	r.GET("/api/philosophies", pyh.List)
	r.GET("/api/philosophies/:id", pyh.Get)

	// Themes — the perennial threads across traditions
	th := handlers.NewThemeHandler(s)
	r.GET("/api/themes", th.List)
	r.GET("/api/themes/:id", th.Get)

	// Evidence — neuroscience & neuropsychology
	eh := handlers.NewEvidenceHandler(s)
	r.GET("/api/evidence", eh.List)
	r.GET("/api/evidence/:id", eh.Get)

	// --- HTML Pages (HTMX + Tailwind) ---

	pages := handlers.NewPages(q, tmpl)

	r.GET("/", pages.Home)
	r.GET("/partials/random-quote", pages.RandomQuotePartial)

	r.GET("/pages/quotes", pages.Quotes)
	r.GET("/pages/philosophers", pages.Philosophers)
	r.GET("/pages/philosophers/:id", pages.PhilosopherDetail)
	r.GET("/pages/philosophies", pages.Philosophies)
	r.GET("/pages/philosophies/:id", pages.PhilosophyDetail)
	r.GET("/pages/themes", pages.Themes)
	r.GET("/pages/themes/:id", pages.ThemeDetail)
	r.GET("/pages/evidence", pages.Evidence)
	r.GET("/pages/evidence/:id", pages.EvidenceDetail)

	return r
}
