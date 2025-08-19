package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/harsinigovindarao/braindump-cli/internal/models"
)

var AllThoughts = []models.Thought{}

const dataFile = "thoughts.jsonl" // Using .jsonl = JSON Lines format

// ✅ Save each thought to memory + file (one JSON per line)
func SaveThought(t models.Thought) {
	AllThoughts = append(AllThoughts, t)

	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Failed to write to file:", err)
		return
	}
	defer file.Close()

	entry, _ := json.Marshal(t)
	file.WriteString(string(entry) + "\n")
}

// ✅ Export all thoughts into a pretty JSON array
func ExportToFile() {
	thoughts := loadThoughtsFromFile()

	file, _ := json.MarshalIndent(thoughts, "", "  ")
	_ = os.WriteFile("thoughts_export.json", file, 0644)
	fmt.Println("💾 Thoughts exported to thoughts_export.json")
}

// ✅ Read all saved thoughts and summarize by category
func SummarizeThoughts(detailed bool) {
	thoughts := loadThoughtsFromFile()
	if len(thoughts) == 0 {
		fmt.Println("📭 No thoughts recorded yet.")
		return
	}

	summary := make(map[string]int)
	for _, t := range thoughts {
		summary[t.Category]++
	}

	fmt.Println("📊 Thought Summary:")
	for cat, count := range summary {
		fmt.Printf(" - %s: %d\n", cat, count)
	}

	if detailed {
		fmt.Println("\n🧠 All Thoughts:")
		sort.SliceStable(thoughts, func(i, j int) bool {
			return thoughts[i].Timestamp.Before(thoughts[j].Timestamp)
		})
		for _, t := range thoughts {
			fmt.Printf(" • [%s][%s] %s — %s\n", t.Category, t.Tone, t.Text, t.Timestamp.Format("Jan 2 15:04"))
		}
	}
}

// 🔧 Utility: Load all thoughts from the file
func loadThoughtsFromFile() []models.Thought {
	var thoughts []models.Thought

	file, err := os.Open(dataFile)
	if err != nil {
		return thoughts
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var t models.Thought
		if err := json.Unmarshal(scanner.Bytes(), &t); err == nil {
			thoughts = append(thoughts, t)
		}
	}

	return thoughts
}

// Exposed for other packages to use
func LoadThoughts() []models.Thought {
	return loadThoughtsFromFile()
}
