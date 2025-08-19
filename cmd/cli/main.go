package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/harsinigovindarao/braindump-cli/internal/app/workers"
	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/prompts"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"

	classification "github.com/harsinigovindarao/braindump-cli/internal/classification"

	"github.com/google/uuid"
)

var ThoughtQueue = make(chan models.Thought, 100)

func main() {
	// Start async processing workers
	go workers.StartProcessor(ThoughtQueue)
	go workers.StartDailyReporter()

	printWelcome()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "":
			continue
		case "/exit":
			fmt.Println("👋 Exiting BrainDump...")
			return
		case "/prompt":
			handlePrompt()
		case "/summary":
			storage.SummarizeThoughts(false)
		case "/summary detail":
			storage.SummarizeThoughts(true)
		case "/export":
			storage.ExportToFile()
		case "/report":
			workers.RunImmediateReport()
		default:
			handleThought(input)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func printWelcome() {
	fmt.Println("🌙 Welcome to 🧠 BrainDump CLI")
	fmt.Println("-------------------------------------------------")
	fmt.Println("Type your thought, or use one of these commands:")
	fmt.Println("   /prompt         → Ask me something")
	fmt.Println("   /summary        → View thought summary")
	fmt.Println("   /summary detail → Full list of all thoughts")
	fmt.Println("   /export         → Save thoughts to a file")
	fmt.Println("   /report         → View top priorities")
	fmt.Println("   /exit           → Quit the app")
	fmt.Println("-------------------------------------------------")
}

func handleThought(input string) {
	history := storage.LoadThoughts()
	if rec := classification.Recommend(input, history); rec != nil {
		fmt.Println("🔁 Similar thought you had earlier:")
		fmt.Println("   🧠", rec.Text)
		fmt.Println("   📅", rec.Timestamp.Format("Jan 2 15:04"))
	}

	t := models.Thought{
		ID:        uuid.New().String(),
		Text:      input,
		Timestamp: time.Now(),
	}
	t.Priority = classification.ScorePriority(t, history)
	ThoughtQueue <- t
}

func handlePrompt() {
	t := prompts.AskAndCapturePrompt()
	t.ID = uuid.New().String()
	t.Priority = classification.ScorePriority(t, storage.LoadThoughts())
	ThoughtQueue <- t
}
