package store

import "perennial-wisdom/models"

// SeedPhilosophers returns wisdom teachers across traditions.
// Epictetus is the apex; each teacher is a lens on the same perennial truth.
func SeedPhilosophers() []models.Philosopher {
	return []models.Philosopher{
		{
			ID: "epictetus", Name: "Epictetus",
			PhilosophyID: "stoic", Era: "50–135 CE",
			Bio: "Born a slave in Hierapolis. His master Epaphroditus broke his leg; Epictetus reportedly said 'I told you it would break.' Freed and taught in Nicopolis. Never wrote — his student Arrian recorded his Discourses and the Enchiridion.",
			KeyTeachings: []string{"Dichotomy of control", "Prohairesis (moral choice)", "Role ethics", "The discipline of desire, action, and assent"},
		},
		{
			ID: "marcus-aurelius", Name: "Marcus Aurelius",
			PhilosophyID: "stoic", Era: "121–180 CE",
			Bio: "Roman Emperor and philosopher-king. His private journal, Meditations, was never meant for publication — raw Stoic self-talk written during military campaigns on the Germanic frontier.",
			KeyTeachings: []string{"Memento mori", "Cosmopolitanism", "Impermanence of all things", "The inner citadel"},
		},
		{
			ID: "seneca", Name: "Seneca",
			PhilosophyID: "stoic", Era: "4 BCE–65 CE",
			Bio: "Roman statesman, dramatist, tutor to Nero. His Letters to Lucilius are among the most practical philosophical texts ever written. Forced to commit suicide by Nero; faced it with Stoic composure.",
			KeyTeachings: []string{"Shortness of life", "Premeditatio malorum", "Anger as temporary madness", "Voluntary discomfort"},
		},
		{
			ID: "epicurus", Name: "Epicurus",
			PhilosophyID: "epicurean", Era: "341–270 BCE",
			Bio: "Founded the Garden — a philosophical community open to women and slaves. Taught that philosophy's purpose is to alleviate suffering. Lived on bread, water, and cheese. 'Send me some cheese, that I may have a feast.'",
			KeyTeachings: []string{"Tetrapharmakos (four-part remedy)", "Death is nothing to us", "Friendship as highest pleasure", "Hierarchy of desires"},
		},
		{
			ID: "diogenes", Name: "Diogenes of Sinope",
			PhilosophyID: "cynic", Era: "412–323 BCE",
			Bio: "Lived in a ceramic jar in the Athenian agora. Carried a lantern in daylight 'looking for an honest man.' When Alexander the Great offered him anything, he said 'Stand out of my sunlight.' The original punk philosopher.",
			KeyTeachings: []string{"Radical self-sufficiency", "Cosmopolitanism", "Defiance of convention", "Living according to nature"},
		},
		{
			ID: "socrates", Name: "Socrates",
			PhilosophyID: "socratic", Era: "470–399 BCE",
			Bio: "Wrote nothing. Taught by questioning. Convicted of corrupting the youth of Athens and impiety. Drank hemlock serenely. His method — elenchus — remains the most powerful tool for exposing assumptions.",
			KeyTeachings: []string{"Socratic method", "Care of the soul", "Virtue as knowledge", "Socratic ignorance"},
		},
		{
			ID: "buddha", Name: "Siddhartha Gautama (The Buddha)",
			PhilosophyID: "buddhist", Era: "563–483 BCE",
			Bio: "Prince who renounced wealth after encountering old age, sickness, and death. Attained enlightenment under the Bodhi tree. Taught the Middle Way for 45 years. 'Be a lamp unto yourself.'",
			KeyTeachings: []string{"Four Noble Truths", "Eightfold Path", "Dependent origination", "Non-self (anatta)"},
		},
		{
			ID: "rumi", Name: "Jalal ad-Din Rumi",
			PhilosophyID: "sufi", Era: "1207–1273 CE",
			Bio: "Persian poet and Sufi mystic. His encounter with the wandering dervish Shams-i-Tabrizi transformed him from scholar to ecstatic poet. His Masnavi is called 'the Quran in Persian.' Founded the Mevlevi (Whirling Dervish) order.",
			KeyTeachings: []string{"Divine love as the supreme reality", "The wound is where the light enters", "Ego-death (fana)", "Unity of being"},
		},
		{
			ID: "shankara", Name: "Adi Shankara",
			PhilosophyID: "vedantic", Era: "788–820 CE",
			Bio: "Indian philosopher who consolidated Advaita (non-dual) Vedanta. Traveled across India debating scholars. Established four monastic centers (mathas). Died at 32, having reshaped Indian philosophy permanently.",
			KeyTeachings: []string{"Brahman alone is real; the world is appearance (maya)", "Self-inquiry", "Neti neti — not this, not this", "Liberation through knowledge (jnana)"},
		},
		{
			ID: "laozi", Name: "Lao Tzu",
			PhilosophyID: "taoist", Era: "6th century BCE (traditional)",
			Bio: "Semi-legendary sage, said to have been an archivist of the Zhou court. Reportedly wrote the Tao Te Ching at a border crossing while leaving civilization. The text's 5,000 characters have influenced millennia of thought.",
			KeyTeachings: []string{"Wu wei (non-action)", "The Tao that can be named is not the eternal Tao", "Water as the supreme metaphor", "Return to simplicity"},
		},
		{
			ID: "krishnamurti", Name: "Jiddu Krishnamurti",
			PhilosophyID: "krishnamurti", Era: "1895–1986 CE",
			Bio: "Raised by the Theosophical Society as the expected World Teacher. At 34, dissolved the Order of the Star, declaring 'Truth is a pathless land.' Spent 60 years giving talks worldwide, refusing followers, methods, and authority.",
			KeyTeachings: []string{"Freedom from the known", "The observer is the observed", "Choiceless awareness", "Thought as time and sorrow"},
		},
	}
}
