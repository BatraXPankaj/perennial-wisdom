package store

import "perennial-wisdom/models"

// SeedThemes returns universal wisdom threads that weave across traditions.
// These are the "perennial" — the same truths rediscovered independently.
func SeedThemes() []models.Theme {
	return []models.Theme{
		{
			ID: "control", Name: "Dichotomy of Control",
			Description: "Distinguishing what is within our power (judgments, intentions, desires) from what is not (external events, others' actions, the body). The root insight of Stoicism, mirrored in Buddhism's acceptance and Taoism's wu wei.",
			PhilosophyIDs: []string{"stoic", "buddhist", "taoist", "krishnamurti"},
		},
		{
			ID: "impermanence", Name: "Impermanence",
			Description: "Nothing lasts. The Buddhist anicca, Stoic memento mori, Heraclitan flux. Neuroscience confirms: the brain itself is in constant structural change (neuroplasticity). Clinging to permanence is the root of suffering.",
			PhilosophyIDs: []string{"stoic", "buddhist", "taoist", "vedantic"},
		},
		{
			ID: "detachment", Name: "Detachment & Non-Attachment",
			Description: "Not indifference, but freedom from compulsive clinging. Stoic apatheia, Buddhist upekkha (equanimity), Vedantic vairagya, Taoist wu wei. The common thread: suffering arises from grasping, not from events.",
			PhilosophyIDs: []string{"stoic", "buddhist", "vedantic", "taoist", "cynic", "krishnamurti"},
		},
		{
			ID: "self-inquiry", Name: "Self-Inquiry & Examination",
			Description: "Socrates' 'know thyself,' Vedanta's 'Who am I?' (Atma Vichara), Krishnamurti's 'the observer is the observed,' Buddhist vipassana. The examined life as the only life worth living.",
			PhilosophyIDs: []string{"socratic", "vedantic", "buddhist", "krishnamurti", "stoic"},
		},
		{
			ID: "virtue", Name: "Virtue as the Highest Good",
			Description: "Stoic arete, Socratic virtue-as-knowledge, Buddhist sila, the Cynic's radical moral life. Not rule-following but alignment with one's deepest nature. Neuroscience links prosocial behavior to dopamine and oxytocin reward circuits.",
			PhilosophyIDs: []string{"stoic", "socratic", "cynic", "buddhist"},
		},
		{
			ID: "ego-dissolution", Name: "Ego Dissolution",
			Description: "Sufi fana, Buddhist anatta, Vedantic 'Atman is Brahman,' Krishnamurti's 'freedom from the known.' The perennial insight: the separate self is a construct. fMRI studies show default mode network quieting during meditation — the neural correlate of ego dissolution.",
			PhilosophyIDs: []string{"sufi", "buddhist", "vedantic", "krishnamurti", "taoist"},
		},
		{
			ID: "present-moment", Name: "Present-Moment Awareness",
			Description: "Stoic prosoche (attention), Buddhist sati (mindfulness), Krishnamurti's choiceless awareness. The past is memory, the future is imagination — only the present is real. Neuroscience: mindfulness thickens the prefrontal cortex and reduces amygdala reactivity.",
			PhilosophyIDs: []string{"stoic", "buddhist", "krishnamurti", "taoist", "sufi"},
		},
		{
			ID: "suffering", Name: "The Nature of Suffering",
			Description: "Buddhism's dukkha, Stoicism's 'it is not things that disturb us but our judgments about them,' Epicurus' hierarchy of desires. Suffering is not in events but in the mind's relationship to events. Neuropsychology: cognitive reappraisal literally changes neural pain responses.",
			PhilosophyIDs: []string{"buddhist", "stoic", "epicurean", "vedantic", "krishnamurti"},
		},
		{
			ID: "simplicity", Name: "Simplicity & Voluntary Poverty",
			Description: "Cynic asceticism, Epicurus' bread and cheese, Taoist pu (the uncarved block), Stoic voluntary discomfort. Excess creates dependency; simplicity creates freedom. Hedonic adaptation research confirms: more stuff ≠ more satisfaction.",
			PhilosophyIDs: []string{"cynic", "epicurean", "taoist", "stoic", "buddhist"},
		},
		{
			ID: "death", Name: "Contemplation of Death",
			Description: "Stoic memento mori, Epicurus' 'death is nothing to us,' Buddhist maranasati (death meditation), Socrates drinking hemlock serenely. Death awareness sharpens life. Terror Management Theory: conscious death reflection reduces unconscious anxiety and increases meaning.",
			PhilosophyIDs: []string{"stoic", "epicurean", "buddhist", "socratic"},
		},
	}
}
