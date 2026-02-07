package models

// Theme represents a universal wisdom thread that runs across traditions.
// Themes are the connective tissue — the "perennial" in perennial wisdom.
// Examples: impermanence, detachment, virtue, self-inquiry, acceptance.
type Theme struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	PhilosophyIDs []string `json:"philosophy_ids"`
}
