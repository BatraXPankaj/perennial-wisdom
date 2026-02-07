package store

import "perennial-wisdom/models"

// SeedQuotes returns perennial wisdom quotes across traditions.
// Epictetus dominates; others provide cross-tradition resonance.
func SeedQuotes() []models.Quote {
	return []models.Quote{
		// — Epictetus (apex) —
		{
			ID: "e1", Text: "It's not what happens to you, but how you react to it that matters.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Discourses",
			ThemeIDs: []string{"control", "suffering"}, EvidenceIDs: []string{"cognitive-reappraisal"},
		},
		{
			ID: "e2", Text: "Make the best use of what is in your power, and take the rest as it happens.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Discourses",
			ThemeIDs: []string{"control", "detachment"},
		},
		{
			ID: "e3", Text: "Man is not worried by real problems so much as by his imagined anxieties about real problems.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Enchiridion",
			ThemeIDs: []string{"suffering", "control"}, EvidenceIDs: []string{"cognitive-reappraisal"},
		},
		{
			ID: "e4", Text: "Wealth consists not in having great possessions, but in having few wants.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Discourses",
			ThemeIDs: []string{"simplicity", "detachment"}, EvidenceIDs: []string{"hedonic-treadmill"},
		},
		{
			ID: "e5", Text: "First say to yourself what you would be; and then do what you have to do.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Discourses",
			ThemeIDs: []string{"virtue", "self-inquiry"},
		},
		{
			ID: "e6", Text: "No man is free who is not master of himself.",
			PhilosopherID: "epictetus", PhilosophyID: "stoic", Source: "Fragments",
			ThemeIDs: []string{"control", "detachment"},
		},

		// — Marcus Aurelius —
		{
			ID: "ma1", Text: "You have power over your mind — not outside events. Realize this, and you will find strength.",
			PhilosopherID: "marcus-aurelius", PhilosophyID: "stoic", Source: "Meditations",
			ThemeIDs: []string{"control", "suffering"}, EvidenceIDs: []string{"cognitive-reappraisal"},
		},
		{
			ID: "ma2", Text: "Think of yourself as dead. You have lived your life. Now, take what's left and live it properly.",
			PhilosopherID: "marcus-aurelius", PhilosophyID: "stoic", Source: "Meditations",
			ThemeIDs: []string{"death", "present-moment"}, EvidenceIDs: []string{"death-awareness"},
		},
		{
			ID: "ma3", Text: "The universe is change; our life is what our thoughts make it.",
			PhilosopherID: "marcus-aurelius", PhilosophyID: "stoic", Source: "Meditations",
			ThemeIDs: []string{"impermanence", "suffering"},
		},

		// — Seneca —
		{
			ID: "s1", Text: "We suffer more often in imagination than in reality.",
			PhilosopherID: "seneca", PhilosophyID: "stoic", Source: "Letters to Lucilius",
			ThemeIDs: []string{"suffering", "control"}, EvidenceIDs: []string{"cognitive-reappraisal"},
		},
		{
			ID: "s2", Text: "It is not that we have a short time to live, but that we waste a great deal of it.",
			PhilosopherID: "seneca", PhilosophyID: "stoic", Source: "On the Shortness of Life",
			ThemeIDs: []string{"death", "present-moment"},
		},

		// — Epicurus —
		{
			ID: "ep1", Text: "Death does not concern us, because as long as we exist, death is not here. And when it does come, we no longer exist.",
			PhilosopherID: "epicurus", PhilosophyID: "epicurean", Source: "Letter to Menoeceus",
			ThemeIDs: []string{"death", "suffering"},
		},
		{
			ID: "ep2", Text: "Do not spoil what you have by desiring what you have not; remember that what you now have was once among the things you only hoped for.",
			PhilosopherID: "epicurus", PhilosophyID: "epicurean", Source: "Vatican Sayings",
			ThemeIDs: []string{"detachment", "simplicity", "present-moment"}, EvidenceIDs: []string{"hedonic-treadmill"},
		},

		// — Diogenes (Cynic) —
		{
			ID: "d1", Text: "It is the privilege of the gods to want nothing, and of godlike men to want little.",
			PhilosopherID: "diogenes", PhilosophyID: "cynic", Source: "Lives of Eminent Philosophers, Diogenes Laertius",
			ThemeIDs: []string{"simplicity", "detachment"}, EvidenceIDs: []string{"hedonic-treadmill"},
		},

		// — Socrates —
		{
			ID: "so1", Text: "The unexamined life is not worth living.",
			PhilosopherID: "socrates", PhilosophyID: "socratic", Source: "Apology, Plato",
			ThemeIDs: []string{"self-inquiry", "virtue"},
		},
		{
			ID: "so2", Text: "I know that I know nothing.",
			PhilosopherID: "socrates", PhilosophyID: "socratic", Source: "Apology, Plato",
			ThemeIDs: []string{"self-inquiry", "ego-dissolution"},
		},

		// — Buddha —
		{
			ID: "b1", Text: "In the end, only three things matter: how much you loved, how gently you lived, and how gracefully you let go of things not meant for you.",
			PhilosopherID: "buddha", PhilosophyID: "buddhist", Source: "Attributed",
			ThemeIDs: []string{"detachment", "impermanence", "virtue"},
		},
		{
			ID: "b2", Text: "You only lose what you cling to.",
			PhilosopherID: "buddha", PhilosophyID: "buddhist", Source: "Attributed",
			ThemeIDs: []string{"detachment", "suffering"},
		},
		{
			ID: "b3", Text: "Nothing is permanent. Everything is subject to change. Being is always becoming.",
			PhilosopherID: "buddha", PhilosophyID: "buddhist", Source: "Attributed",
			ThemeIDs: []string{"impermanence"}, EvidenceIDs: []string{"mindfulness-cortex"},
		},

		// — Rumi (Sufi) —
		{
			ID: "r1", Text: "The wound is the place where the Light enters you.",
			PhilosopherID: "rumi", PhilosophyID: "sufi", Source: "Collected Poems",
			ThemeIDs: []string{"suffering", "ego-dissolution"},
		},
		{
			ID: "r2", Text: "Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.",
			PhilosopherID: "rumi", PhilosophyID: "sufi", Source: "Collected Poems",
			ThemeIDs: []string{"self-inquiry", "control"},
		},

		// — Adi Shankara (Vedanta) —
		{
			ID: "sh1", Text: "Brahman alone is real; the world is appearance. The self is nothing but Brahman.",
			PhilosopherID: "shankara", PhilosophyID: "vedantic", Source: "Vivekachudamani",
			ThemeIDs: []string{"ego-dissolution", "self-inquiry"}, EvidenceIDs: []string{"self-referential-processing"},
		},

		// — Lao Tzu (Tao) —
		{
			ID: "lt1", Text: "Nature does not hurry, yet everything is accomplished.",
			PhilosopherID: "laozi", PhilosophyID: "taoist", Source: "Tao Te Ching",
			ThemeIDs: []string{"detachment", "present-moment"},
		},
		{
			ID: "lt2", Text: "When I let go of what I am, I become what I might be.",
			PhilosopherID: "laozi", PhilosophyID: "taoist", Source: "Tao Te Ching",
			ThemeIDs: []string{"detachment", "ego-dissolution"},
		},

		// — Krishnamurti —
		{
			ID: "k1", Text: "It is no measure of health to be well adjusted to a profoundly sick society.",
			PhilosopherID: "krishnamurti", PhilosophyID: "krishnamurti", Source: "Attributed",
			ThemeIDs: []string{"self-inquiry", "virtue"},
		},
		{
			ID: "k2", Text: "The ability to observe without evaluating is the highest form of intelligence.",
			PhilosopherID: "krishnamurti", PhilosophyID: "krishnamurti", Source: "Freedom from the Known",
			ThemeIDs: []string{"present-moment", "ego-dissolution"}, EvidenceIDs: []string{"dmn-meditation"},
		},
	}
}
