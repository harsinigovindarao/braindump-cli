package models

import "time"

type Thought struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Category  string    `json:"category"`
	Tone      string    `json:"tone"`
	Timestamp time.Time `json:"timestamp"`
	Prompt    string    `json:"prompt"`
	Priority  int       `json:"priority"`
}
