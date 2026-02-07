package models

// Philosopher represents a wisdom teacher across traditions.
// Epictetus is the apex; others provide cross-tradition resonance.
type Philosopher struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	PhilosophyID string   `json:"philosophy_id"`
	Era          string   `json:"era"`
	Bio          string   `json:"bio"`
	KeyTeachings []string `json:"key_teachings"`
}
