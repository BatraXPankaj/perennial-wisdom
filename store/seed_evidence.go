package store

import "perennial-wisdom/models"

// SeedEvidence returns neuroscience and neuropsychology findings
// that empirically validate ancient contemplative insights.
func SeedEvidence() []models.Evidence {
	return []models.Evidence{
		{
			ID: "dmn-meditation", Title: "Default Mode Network & Meditation",
			Finding: "Experienced meditators show reduced activity in the default mode network (DMN) — the brain's 'selfing' circuit. The DMN generates the narrative self, mind-wandering, and rumination. Its quieting maps directly to the contemplative experience of ego dissolution reported across Sufi, Buddhist, and Vedantic traditions.",
			Field: "neuroscience", Source: "Brewer et al., PNAS, 2011",
			ThemeIDs: []string{"ego-dissolution", "present-moment"},
		},
		{
			ID: "cognitive-reappraisal", Title: "Cognitive Reappraisal Changes Neural Pain Responses",
			Finding: "Reframing the meaning of an event (cognitive reappraisal) reduces activation in the amygdala and increases prefrontal cortex engagement. This is the neural mechanism behind Epictetus' core teaching: 'It is not things that disturb us, but our judgments about things.' Stoic practice is, neurologically, a reappraisal protocol.",
			Field: "neuropsychology", Source: "Ochsner & Gross, Trends in Cognitive Sciences, 2005",
			ThemeIDs: []string{"suffering", "control"},
		},
		{
			ID: "mindfulness-cortex", Title: "Mindfulness Thickens the Prefrontal Cortex",
			Finding: "8 weeks of mindfulness meditation (MBSR) increases cortical thickness in the prefrontal cortex (executive function, attention) and reduces grey matter in the amygdala (fear, reactivity). Ancient practice, measurable structural brain change.",
			Field: "neuroscience", Source: "Hölzel et al., Psychiatry Research: Neuroimaging, 2011",
			ThemeIDs: []string{"present-moment", "detachment"},
		},
		{
			ID: "hedonic-treadmill", Title: "Hedonic Adaptation and the Simplicity Insight",
			Finding: "Lottery winners return to baseline happiness within months. The hedonic treadmill confirms what Epicurus, Diogenes, and Seneca taught: external acquisitions produce diminishing returns. Lasting well-being comes from internal states, not circumstances.",
			Field: "psychology", Source: "Brickman & Campbell, 1971; updated by Diener et al.",
			ThemeIDs: []string{"simplicity", "detachment", "suffering"},
		},
		{
			ID: "death-awareness", Title: "Terror Management & Death Contemplation",
			Finding: "Terror Management Theory shows that unconscious death anxiety drives materialism, tribalism, and aggression. But conscious, deliberate death reflection (Stoic memento mori, Buddhist maranasati) has the opposite effect: it increases gratitude, prosocial behavior, and meaning-making. The ancients were right — the direction matters.",
			Field: "psychology", Source: "Cozzolino et al., Personality and Social Psychology Bulletin, 2004",
			ThemeIDs: []string{"death", "present-moment", "virtue"},
		},
		{
			ID: "self-referential-processing", Title: "The Constructed Self — Neuroscience of Anatta",
			Finding: "The brain has no single 'self center.' Self-referential processing is distributed across the DMN, medial prefrontal cortex, and posterior cingulate. The 'self' is a process, not a thing — confirming Buddhist anatta, Vedantic maya, and Krishnamurti's 'the observer is the observed.'",
			Field: "neuroscience", Source: "Northoff et al., Neuroscience & Biobehavioral Reviews, 2006",
			ThemeIDs: []string{"ego-dissolution", "self-inquiry"},
		},
	}
}
