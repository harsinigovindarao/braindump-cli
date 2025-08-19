package workers

import (
	"fmt"
	"sort"
	"time"

	"github.com/harsinigovindarao/braindump-cli/internal/storage"
)

// StartDailyReporter runs periodically (every 5s for demo, every 24h in real use)
func StartDailyReporter() {
	interval := 50 * time.Second // 🧪 Demo mode: check every 50 seconds
	// interval := 24 * time.Hour // 🕛 Real use: check once daily

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			printTopThoughts("🧠 Top Priority Thoughts:")
			fmt.Printf("⏳ Next check in %s...\n", interval)
		}
	}()
}

// RunImmediateReport runs once and prints top 3 thoughts immediately
func RunImmediateReport() {
	printTopThoughts("🧠 Immediate Thought Report:")
}

// printTopThoughts loads and prints top 3 thoughts by priority
func printTopThoughts(header string) {
	thoughts := storage.LoadThoughts()
	if len(thoughts) == 0 {
		fmt.Println("📭 No thoughts to report.")
		return
	}

	sort.SliceStable(thoughts, func(i, j int) bool {
		return thoughts[i].Priority > thoughts[j].Priority
	})

	fmt.Println(header)
	for i := 0; i < len(thoughts) && i < 3; i++ {
		t := thoughts[i]
		fmt.Printf(" - [%s] %s (Priority: %d)\n", t.Tone, t.Text, t.Priority)
	}
}
