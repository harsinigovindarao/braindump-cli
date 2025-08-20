package prompts

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"
	"github.com/harsinigovindarao/braindump-cli/utils"
)

var samplePrompts = []string{
	"What made you smile recently?",
	"What’s a small win you had today?",
	"What would you do if you had no fear?",
	"What’s something you want to learn?",
	"What’s on your mind right now?",
	"What challenge are you currently facing?",
}

func GetRandomPrompt() string {
	return samplePrompts[rand.Intn(len(samplePrompts))]
}

func AskAndCapturePrompt() models.Thought {
	prompt := GetRandomPrompt()
	fmt.Println("💡 Prompt:", prompt)
	fmt.Print("📝 Your answer: ")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)

	// Build the current thought
	currentThought := models.Thought{
		ID:        uuid.NewString(), // fixed: use UUID
		Text:      response,
		Timestamp: time.Now(),
		Prompt:    prompt,
	}

	// Recommend similar thoughts
	history := storage.LoadThoughts()
	recs := utils.Recommend(currentThought, history) // slice of thoughts

	if len(recs) > 0 {
		fmt.Println("🔁 Similar thoughts you had earlier:")
		for i, r := range recs {
			if i >= 3 { // limit to top 3 for CLI
				break
			}
			fmt.Println("   🧠", r.Text)
			fmt.Println("   📅", r.Timestamp.Format("Jan 2 15:04"))
		}
	}

	return currentThought
}
