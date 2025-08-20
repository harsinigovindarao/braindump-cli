// utils.go
package utils

import (
    "github.com/harsinigovindarao/braindump-cli/internal/models"
)

// Classify classifies the input text into categories
func Classify(text string) string {
    // TODO: Replace with real logic
    return "general"
}

// DetectTone detects emotional tone in the input
func DetectTone(text string) string {
    // TODO: Replace with real logic
    return "neutral"
}

// ScorePriority assigns a priority score to text
func ScorePriority(thought models.Thought, history []models.Thought) int {
    // use thought.Text, thought.Category, history etc.
    return 1
}

// Recommend gives a recommendation based on text
func Recommend(thought models.Thought, allThoughts []models.Thought) []models.Thought {
    var recommendations []models.Thought
    for _, t := range allThoughts {
        if t.ID != thought.ID && (t.Category == thought.Category || t.Tone == thought.Tone) {
            recommendations = append(recommendations, t)
        }
    }
    return recommendations
}
