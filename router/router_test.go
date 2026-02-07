package router_test

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	"perennial-wisdom/db"
	"perennial-wisdom/router"
	"perennial-wisdom/store"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// setupTestRouter creates a fully-wired router with in-memory DB and minimal templates.
func setupTestRouter(t *testing.T) *gin.Engine {
	t.Helper()

	// In-memory store for JSON API
	s := store.New()

	// In-memory SQLite for HTML pages
	conn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	conn.Exec("PRAGMA foreign_keys=ON")
	database := sqlx.NewDb(conn, "sqlite")

	db.Migrate(database)
	db.Seed(database)

	q := db.NewQueries(database)

	// Minimal template set for testing — each named template produces predictable output
	tmpl := template.Must(template.New("base").Parse(`{{define "base"}}<!DOCTYPE html><title>{{.Title}}</title>{{end}}`))
	template.Must(tmpl.New("random-quote").Parse(`{{define "random-quote"}}<q>{{.Text}}</q>{{end}}`))

	return router.Setup(s, q, tmpl)
}

// ---- Route Existence ----

func TestRouteHealth(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("/health: expected 200, got %d", w.Code)
	}
}

// ---- JSON API Integration ----

func TestAPIQuotesIntegration(t *testing.T) {
	r := setupTestRouter(t)

	tests := []struct {
		path       string
		expectCode int
		checkBody  func(t *testing.T, body map[string]interface{})
	}{
		{"/api/quotes", 200, func(t *testing.T, body map[string]interface{}) {
			count := int(body["count"].(float64))
			if count == 0 {
				t.Error("expected quotes, got 0")
			}
		}},
		{"/api/quotes/random", 200, func(t *testing.T, body map[string]interface{}) {
			if body["philosopher"] == nil {
				t.Error("expected philosopher in response")
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.expectCode {
				t.Errorf("%s: expected %d, got %d", tt.path, tt.expectCode, w.Code)
			}

			var body map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &body)
			tt.checkBody(t, body)
		})
	}
}

func TestAPIPhilosophersIntegration(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophers", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected philosophers, got 0")
	}
}

func TestAPIPhilosophiesIntegration(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophies", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected philosophies, got 0")
	}
}

func TestAPIThemesIntegration(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/themes", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected themes, got 0")
	}
}

func TestAPIEvidenceIntegration(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/evidence", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected evidence, got 0")
	}
}

// ---- HTML Pages Integration ----

func TestPageHome(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("/: expected 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("expected text/html, got %s", contentType)
	}
}

func TestPageRandomQuotePartial(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/partials/random-quote", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("/partials/random-quote: expected 200, got %d", w.Code)
	}
}

// ---- 404 for unknown routes ----

func TestNotFoundRoute(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 for unknown route, got %d", w.Code)
	}
}

// ---- Method Not Allowed ----

func TestMethodNotAllowed(t *testing.T) {
	r := setupTestRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/quotes", nil)
	r.ServeHTTP(w, req)

	// Gin returns 404 for POST on GET-only routes by default
	if w.Code == http.StatusOK {
		t.Errorf("expected non-200 for POST on GET route, got %d", w.Code)
	}
}
