package utils

import "strings"

func Classify(text string) string {
	text = strings.ToLower(text)
	switch {
	case strings.Contains(text, "buy"), strings.Contains(text, "get"):
		return "ToDo"
	case strings.Contains(text, "idea"), strings.Contains(text, "what if"):
		return "Idea"
	case strings.Contains(text, "remember"), strings.Contains(text, "remind"):
		return "Reminder"
	case strings.Contains(text, "feel"), strings.Contains(text, "why"):
		return "Journal"
	default:
		return "Unsorted"
	}
}

func DetectTone(text string) string {
	text = strings.ToLower(text)

	// Clean informal phrases
	text = strings.ReplaceAll(text, "wanna", "want to")
	text = strings.ReplaceAll(text, "gonna", "going to")

	negativeWords := []string{
		"kill", "hate", "angry", "sad", "depressed", "upset", "mad", "annoyed",
		"rage", "cry", "hurt", "pain", "die", "murder", "revenge", "destroy", "suicide", "attack",
	}
	positiveWords := []string{
		"happy", "excited", "grateful", "joy", "love", "awesome", "fun", "peace", "calm",
		"great", "wonderful", "thankful", "cheerful", "enjoy", "smile",
	}

	for _, word := range negativeWords {
		if strings.Contains(text, word) {
			return "Negative"
		}
	}
	for _, word := range positiveWords {
		if strings.Contains(text, word) {
			return "Positive"
		}
	}

	return "Neutral"
}
