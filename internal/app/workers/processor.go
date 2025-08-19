package workers

import (
	"fmt"
	"strings"

	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"
)

func StartProcessor(ch <-chan models.Thought) {
	go func() {
		for t := range ch {
			// ✏️ Auto classify thought
			t.Category = utils.Classify(t.Text)

			// 🎭 Detect tone
			t.Tone = utils.DetectTone(t.Text)

			// 🔢 Assign basic priority
			t.Priority = assignPriority(t.Text, t.Tone)

			// 🧠 Output to user
			fmt.Println("🧠 Categorized as:", t.Category)
			fmt.Println("🧘 Tone:", t.Tone)

			// 💾 Save to storage
			storage.SaveThought(t)
		}
	}()
}

// assignPriority sets a basic priority score (0 to 5)
func assignPriority(text string, tone string) int {
	text = strings.ToLower(text)

	switch {
	case strings.Contains(text, "urgent"), strings.Contains(text, "important"), strings.Contains(text, "now"):
		return 5
	case strings.Contains(text, "remind"), strings.Contains(text, "soon"):
		return 4
	case tone == "Negative":
		return 3
	case tone == "Neutral":
		return 2
	default:
		return 1
	}
}
