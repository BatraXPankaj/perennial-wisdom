package store

import "perennial-wisdom/models"

// SeedPhilosophies returns the schools of perennial wisdom.
// Stoic is primary; others reveal the shared substrate of human insight.
func SeedPhilosophies() []models.Philosophy {
	return []models.Philosophy{
		{
			ID:   "stoic",
			Name: "Stoicism",
			Origin: "Greece, 3rd century BCE — Zeno of Citium",
			CorePrinciples: []string{
				"Dichotomy of control: distinguish what is up to us from what is not",
				"Virtue (arete) is the sole good",
				"Live according to nature and reason (logos)",
				"Negative visualization (premeditatio malorum)",
				"The obstacle is the way",
			},
			RelatedIDs: []string{"buddhist", "vedantic", "taoist"},
		},
		{
			ID:   "epicurean",
			Name: "Epicureanism",
			Origin: "Greece, 4th century BCE — Epicurus",
			CorePrinciples: []string{
				"Pleasure (ataraxia — tranquility) is the highest good",
				"Absence of pain (aponia) over active pleasure",
				"Fear of death is irrational",
				"Simple living and friendship are supreme",
				"Natural and necessary desires vs vain desires",
			},
			RelatedIDs: []string{"stoic", "buddhist"},
		},
		{
			ID:   "cynic",
			Name: "Cynicism",
			Origin: "Greece, 5th century BCE — Antisthenes, Diogenes of Sinope",
			CorePrinciples: []string{
				"Virtue is the only good",
				"Reject all conventional desires: wealth, power, fame",
				"Live in agreement with nature",
				"Radical freedom through radical simplicity",
				"Parrhesia — fearless, shameless speech",
			},
			RelatedIDs: []string{"stoic", "taoist", "buddhist"},
		},
		{
			ID:   "socratic",
			Name: "Socratic Philosophy",
			Origin: "Greece, 5th century BCE — Socrates",
			CorePrinciples: []string{
				"The unexamined life is not worth living",
				"I know that I know nothing (Socratic ignorance)",
				"Virtue is knowledge; no one errs willingly",
				"Dialectic questioning as path to truth",
				"Care of the soul above all else",
			},
			RelatedIDs: []string{"stoic", "vedantic"},
		},
		{
			ID:   "buddhist",
			Name: "Buddhism",
			Origin: "India, 5th century BCE — Siddhartha Gautama",
			CorePrinciples: []string{
				"Four Noble Truths: suffering, its cause, its end, the path",
				"Impermanence (anicca) of all phenomena",
				"Non-self (anatta) — no fixed, separate self",
				"The Middle Way between asceticism and indulgence",
				"Mindfulness and meditation as liberation",
			},
			RelatedIDs: []string{"stoic", "vedantic", "taoist"},
		},
		{
			ID:   "sufi",
			Name: "Sufism",
			Origin: "Islamic world, 8th century CE — mystical tradition within Islam",
			CorePrinciples: []string{
				"Fana — annihilation of the ego-self in the Divine",
				"Divine love as the supreme path",
				"The inner (batin) meaning beneath the outer (zahir)",
				"Zikr — remembrance of God through repetition",
				"The heart as the organ of spiritual perception",
			},
			RelatedIDs: []string{"vedantic", "buddhist"},
		},
		{
			ID:   "vedantic",
			Name: "Vedanta",
			Origin: "India, ~800 BCE onward — Upanishads, Shankara",
			CorePrinciples: []string{
				"Atman (self) is Brahman (ultimate reality)",
				"The world of multiplicity is maya (illusion)",
				"Self-inquiry (Atma Vichara): Who am I?",
				"Liberation (moksha) through knowledge (jnana)",
				"Neti neti — not this, not this — via negativa",
			},
			RelatedIDs: []string{"buddhist", "sufi", "stoic"},
		},
		{
			ID:   "taoist",
			Name: "Taoism",
			Origin: "China, 6th century BCE — Lao Tzu",
			CorePrinciples: []string{
				"The Tao that can be named is not the eternal Tao",
				"Wu wei — effortless action, non-forcing",
				"Harmony with the natural flow of things",
				"Simplicity (pu) — the uncarved block",
				"Yin-yang: complementary opposites",
			},
			RelatedIDs: []string{"stoic", "buddhist", "cynic"},
		},
		{
			ID:   "krishnamurti",
			Name: "Krishnamurti's Teaching",
			Origin: "India/Global, 20th century — Jiddu Krishnamurti",
			CorePrinciples: []string{
				"Truth is a pathless land — no guru, no method",
				"The observer is the observed",
				"Freedom from the known",
				"Thought is the root of psychological suffering",
				"Choiceless awareness — observation without the observer",
			},
			RelatedIDs: []string{"buddhist", "vedantic", "taoist"},
		},
	}
}
