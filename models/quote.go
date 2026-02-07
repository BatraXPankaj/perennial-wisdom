package models

// Quote is the central entity — a piece of perennial wisdom.
// It links a philosopher's words to themes, a school of thought,
// and optionally to scientific evidence supporting the insight.
type Quote struct {
	ID           string   `json:"id"`
	Text         string   `json:"text"`
	PhilosopherID string  `json:"philosopher_id"`
	PhilosophyID string   `json:"philosophy_id"`
	Source       string   `json:"source"`
	ThemeIDs     []string `json:"theme_ids"`
	EvidenceIDs  []string `json:"evidence_ids,omitempty"`
}
