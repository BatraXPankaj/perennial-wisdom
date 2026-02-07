package models

// Evidence represents scientific research — primarily neuroscience and
// neuropsychology — that validates ancient wisdom empirically.
// Bridges the contemplative and the empirical.
type Evidence struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Finding     string   `json:"finding"`
	Field       string   `json:"field"` // "neuroscience", "neuropsychology", "psychology"
	Source      string   `json:"source"`
	ThemeIDs    []string `json:"theme_ids"`
}
