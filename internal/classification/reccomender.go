package utils

import (
	"braindump-cli/internal/models"
	"strings"
)

// 🔍 Stop words list (you can expand it)
var stopWords = map[string]bool{
	"i": true, "am": true, "is": true, "are": true, "the": true,
	"and": true, "to": true, "a": true, "in": true, "on": true,
	"it": true, "of": true, "for": true, "was": true, "with": true,
}

// Recommend returns the most similar thought from history
func Recommend(input string, history []models.Thought) *models.Thought {
	inputWords := getWords(input)
	var bestMatch *models.Thought
	maxScore := 0

	for _, past := range history {
		score := similarityScore(inputWords, getWords(past.Text))
		if score > maxScore && !strings.EqualFold(past.Text, input) {
			bestMatch = &past
			maxScore = score
		}
	}

	return bestMatch
}

// ✅ Use improved filtering
func getWords(text string) []string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	words := strings.Fields(text)

	// Filter stop words
	var filtered []string
	for _, word := range words {
		if !stopWords[word] {
			filtered = append(filtered, word)
		}
	}
	return filtered
}

// 🔁 Returns how many words are common between two sets
func similarityScore(a, b []string) int {
	score := 0
	wordSet := make(map[string]bool)
	for _, word := range b {
		wordSet[word] = true
	}
	for _, word := range a {
		if wordSet[word] {
			score++
		}
	}
	return score
}
