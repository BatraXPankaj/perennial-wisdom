package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"perennial-wisdom/handlers"
	"perennial-wisdom/models"
	"perennial-wisdom/store"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// testStore returns a minimal Store pre-loaded with a known data set.
func testStore() *store.Store {
	s := &store.Store{
		Quotes:       make(map[string]models.Quote),
		Philosophers: make(map[string]models.Philosopher),
		Philosophies: make(map[string]models.Philosophy),
		Themes:       make(map[string]models.Theme),
		Evidence:     make(map[string]models.Evidence),
	}

	s.Philosophies["stoic"] = models.Philosophy{
		ID: "stoic", Name: "Stoicism", Origin: "Ancient Greece",
		CorePrinciples: []string{"Virtue", "Acceptance"},
		RelatedIDs:     []string{"buddhist"},
	}
	s.Philosophies["buddhist"] = models.Philosophy{
		ID: "buddhist", Name: "Buddhism", Origin: "Ancient India",
		CorePrinciples: []string{"Four Noble Truths"},
	}

	s.Philosophers["epictetus"] = models.Philosopher{
		ID: "epictetus", Name: "Epictetus", PhilosophyID: "stoic",
		Era: "55–135 CE", Bio: "Former slave turned Stoic teacher",
		KeyTeachings: []string{"Dichotomy of control"},
	}
	s.Philosophers["buddha"] = models.Philosopher{
		ID: "buddha", Name: "Siddhartha Gautama", PhilosophyID: "buddhist",
		Era: "563–483 BCE", Bio: "Founder of Buddhism",
		KeyTeachings: []string{"Middle Way"},
	}

	s.Themes["control"] = models.Theme{
		ID: "control", Name: "Sphere of Control",
		Description:   "Focus energy only on what you can influence",
		PhilosophyIDs: []string{"stoic", "buddhist"},
	}
	s.Themes["impermanence"] = models.Theme{
		ID: "impermanence", Name: "Impermanence",
		Description:   "All things change",
		PhilosophyIDs: []string{"buddhist", "stoic"},
	}

	s.Evidence["neuro-control"] = models.Evidence{
		ID: "neuro-control", Title: "Prefrontal Cortex & Control",
		Finding: "PFC activation during reappraisal",
		Field: "neuroscience", Source: "Davidson 2004",
		ThemeIDs: []string{"control"},
	}

	s.Quotes["q1"] = models.Quote{
		ID: "q1", Text: "It is not things that disturb us, but our judgments about them.",
		PhilosopherID: "epictetus", PhilosophyID: "stoic",
		Source: "Enchiridion", ThemeIDs: []string{"control"},
		EvidenceIDs: []string{"neuro-control"},
	}
	s.Quotes["q2"] = models.Quote{
		ID: "q2", Text: "All conditioned things are impermanent.",
		PhilosopherID: "buddha", PhilosophyID: "buddhist",
		Source: "Dhammapada", ThemeIDs: []string{"impermanence"},
	}

	return s
}

// ---- Health ----

func TestHealth(t *testing.T) {
	r := gin.New()
	r.GET("/health", handlers.Health)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["status"] != "ok" {
		t.Errorf("expected status=ok, got %v", body["status"])
	}
	if body["project"] != "Perennial Wisdom API" {
		t.Errorf("unexpected project name: %v", body["project"])
	}
}

// ---- Quotes ----

func TestQuoteList(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes", qh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 2 {
		t.Errorf("expected 2 quotes, got %d", count)
	}
}

func TestQuoteListFilterByPhilosopher(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes", qh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes?philosopher=epictetus", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 quote for epictetus, got %d", count)
	}
}

func TestQuoteListFilterByPhilosophy(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes", qh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes?philosophy=buddhist", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 buddhist quote, got %d", count)
	}
}

func TestQuoteListFilterByTheme(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes", qh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes?theme=control", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 'control' themed quote, got %d", count)
	}
}

func TestQuoteListFilterNoMatch(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes", qh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes?philosopher=nonexistent", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 0 {
		t.Errorf("expected 0 quotes, got %d", count)
	}
}

func TestQuoteGet(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes/:id", qh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes/q1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["philosopher"] != "Epictetus" {
		t.Errorf("expected philosopher=Epictetus, got %v", body["philosopher"])
	}
	if body["philosophy"] != "Stoicism" {
		t.Errorf("expected philosophy=Stoicism, got %v", body["philosophy"])
	}
}

func TestQuoteGetNotFound(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes/:id", qh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestQuoteRandom(t *testing.T) {
	s := testStore()
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes/random", qh.Random)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes/random", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["philosopher"] == nil {
		t.Error("expected philosopher in random quote response")
	}
}

func TestQuoteRandomEmpty(t *testing.T) {
	s := &store.Store{
		Quotes:       make(map[string]models.Quote),
		Philosophers: make(map[string]models.Philosopher),
		Philosophies: make(map[string]models.Philosophy),
		Themes:       make(map[string]models.Theme),
		Evidence:     make(map[string]models.Evidence),
	}
	qh := handlers.NewQuoteHandler(s)

	r := gin.New()
	r.GET("/api/quotes/random", qh.Random)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/quotes/random", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for empty store, got %d", w.Code)
	}
}

// ---- Philosophers ----

func TestPhilosopherList(t *testing.T) {
	s := testStore()
	ph := handlers.NewPhilosopherHandler(s)

	r := gin.New()
	r.GET("/api/philosophers", ph.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophers", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 2 {
		t.Errorf("expected 2 philosophers, got %d", count)
	}
}

func TestPhilosopherListFilterByPhilosophy(t *testing.T) {
	s := testStore()
	ph := handlers.NewPhilosopherHandler(s)

	r := gin.New()
	r.GET("/api/philosophers", ph.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophers?philosophy=stoic", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 stoic philosopher, got %d", count)
	}
}

func TestPhilosopherGet(t *testing.T) {
	s := testStore()
	ph := handlers.NewPhilosopherHandler(s)

	r := gin.New()
	r.GET("/api/philosophers/:id", ph.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophers/epictetus", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["philosophy"] != "Stoicism" {
		t.Errorf("expected philosophy=Stoicism, got %v", body["philosophy"])
	}
}

func TestPhilosopherGetNotFound(t *testing.T) {
	s := testStore()
	ph := handlers.NewPhilosopherHandler(s)

	r := gin.New()
	r.GET("/api/philosophers/:id", ph.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophers/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// ---- Philosophies ----

func TestPhilosophyList(t *testing.T) {
	s := testStore()
	pyh := handlers.NewPhilosophyHandler(s)

	r := gin.New()
	r.GET("/api/philosophies", pyh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophies", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 2 {
		t.Errorf("expected 2 philosophies, got %d", count)
	}
}

func TestPhilosophyGet(t *testing.T) {
	s := testStore()
	pyh := handlers.NewPhilosophyHandler(s)

	r := gin.New()
	r.GET("/api/philosophies/:id", pyh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophies/stoic", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	// Should include related philosophies and philosophers
	if body["related"] == nil {
		t.Error("expected related philosophies")
	}
	if body["philosophers"] == nil {
		t.Error("expected philosophers list")
	}
}

func TestPhilosophyGetNotFound(t *testing.T) {
	s := testStore()
	pyh := handlers.NewPhilosophyHandler(s)

	r := gin.New()
	r.GET("/api/philosophies/:id", pyh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/philosophies/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// ---- Themes ----

func TestThemeList(t *testing.T) {
	s := testStore()
	th := handlers.NewThemeHandler(s)

	r := gin.New()
	r.GET("/api/themes", th.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/themes", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 2 {
		t.Errorf("expected 2 themes, got %d", count)
	}
}

func TestThemeGet(t *testing.T) {
	s := testStore()
	th := handlers.NewThemeHandler(s)

	r := gin.New()
	r.GET("/api/themes/:id", th.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/themes/control", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["quotes"] == nil {
		t.Error("expected quotes in theme detail")
	}
	if body["evidence"] == nil {
		t.Error("expected evidence in theme detail")
	}
	if body["philosophies"] == nil {
		t.Error("expected philosophies in theme detail")
	}
}

func TestThemeGetNotFound(t *testing.T) {
	s := testStore()
	th := handlers.NewThemeHandler(s)

	r := gin.New()
	r.GET("/api/themes/:id", th.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/themes/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// ---- Evidence ----

func TestEvidenceList(t *testing.T) {
	s := testStore()
	eh := handlers.NewEvidenceHandler(s)

	r := gin.New()
	r.GET("/api/evidence", eh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/evidence", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 evidence, got %d", count)
	}
}

func TestEvidenceListFilterByField(t *testing.T) {
	s := testStore()
	eh := handlers.NewEvidenceHandler(s)

	r := gin.New()
	r.GET("/api/evidence", eh.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/evidence?field=neuroscience", nil)
	r.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	count := int(body["count"].(float64))
	if count != 1 {
		t.Errorf("expected 1 neuroscience evidence, got %d", count)
	}

	// Filter by non-matching field
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/evidence?field=philosophy", nil)
	r.ServeHTTP(w2, req2)

	var body2 map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &body2)

	count2 := int(body2["count"].(float64))
	if count2 != 0 {
		t.Errorf("expected 0, got %d", count2)
	}
}

func TestEvidenceGet(t *testing.T) {
	s := testStore()
	eh := handlers.NewEvidenceHandler(s)

	r := gin.New()
	r.GET("/api/evidence/:id", eh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/evidence/neuro-control", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)

	if body["themes"] == nil {
		t.Error("expected themes in evidence detail")
	}
	if body["quotes"] == nil {
		t.Error("expected quotes in evidence detail")
	}
}

func TestEvidenceGetNotFound(t *testing.T) {
	s := testStore()
	eh := handlers.NewEvidenceHandler(s)

	r := gin.New()
	r.GET("/api/evidence/:id", eh.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/evidence/nonexistent", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
