package models

// Philosophy represents a school of perennial wisdom.
// Stoicism is the primary lens; others reveal shared truths across traditions.
type Philosophy struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Origin          string   `json:"origin"`
	CorePrinciples  []string `json:"core_principles"`
	RelatedIDs      []string `json:"related_ids,omitempty"`
}
