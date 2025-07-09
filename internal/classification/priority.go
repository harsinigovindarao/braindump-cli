package utils

import (
	"braindump-cli/internal/models"
	"strings"
)

// ScorePriority assigns a score based on tone and frequency
func ScorePriority(t models.Thought, history []models.Thought) int {
	score := 0

	// Frequency check: more similar = more important
	count := 0
	for _, past := range history {
		if isSimilar(t.Text, past.Text) {
			count++
		}
	}
	score += count * 2

	// Tone boost
	switch strings.ToLower(t.Tone) {
	case "negative":
		score += 3
	case "neutral":
		score += 1
	case "positive":
		score += 0
	}

	return score
}

func isSimilar(a, b string) bool {
	a = strings.ToLower(strings.TrimSpace(a))
	b = strings.ToLower(strings.TrimSpace(b))
	return strings.Contains(a, b) || strings.Contains(b, a)
}
