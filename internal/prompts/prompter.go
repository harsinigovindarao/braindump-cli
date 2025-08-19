package prompts

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"
	"github.com/harsinigovindarao/braindump-cli/internal/utils"

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

	// Recommend similar thoughts
	history := storage.LoadThoughts()
	if rec := utils.Recommend(response, history); rec != nil {
		fmt.Println("🔁 Similar thought you had earlier:")
		fmt.Println("   🧠", rec.Text)
		fmt.Println("   📅", rec.Timestamp.Format("Jan 2 15:04"))
	}

	// Build Thought
	return models.Thought{
		ID:        "1", // replace with UUID if needed
		Text:      response,
		Timestamp: time.Now(),
		Prompt:    prompt,
	}
}
