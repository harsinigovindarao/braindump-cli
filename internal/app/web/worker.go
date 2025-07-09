package web

import (
	"braindump-cli/internal/models"
	"braindump-cli/internal/storage"
	"braindump-cli/internal/classification"
)

// Channels
var ThoughtInputChan = make(chan models.Thought)
var ProcessedThoughtChan = make(chan models.Thought)

// StartWorker launches goroutines for async thought processing
func StartWorker() {
	// Worker for classification, tone, scoring
	go func() {
		for t := range ThoughtInputChan {
			t.Category = utils.Classify(t.Text)
			t.Tone = utils.DetectTone(t.Text)
			t.Priority = utils.ScorePriority(t, storage.LoadThoughts())
			ProcessedThoughtChan <- t
		}
	}()

	// Worker for saving processed thoughts
	go func() {
		for pt := range ProcessedThoughtChan {
			storage.SaveThought(pt)
		}
	}()
}
