package web

import (
	"github.com/harsinigovindarao/braindump-cli/internal/models"
	"github.com/harsinigovindarao/braindump-cli/internal/prompts"
	"github.com/harsinigovindarao/braindump-cli/internal/storage"

	//"github.com/harsinigovindarao/braindump-cli/internal/classification"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/google/uuid"
)

var cachedReport string // 🧠 shared report data

// ServeWebUI starts the web interface and API endpoints
func ServeWebUI() {
	// Serve static HTML and assets
	http.Handle("/", http.FileServer(http.Dir("./static")))


	// POST /thought — save a new thought
	http.HandleFunc("/thought", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var payload struct{ Text string }
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.Text == "" {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}
		StartWorker() // 🔥 start goroutine workers

		t := models.Thought{
			ID:        uuid.NewString(),
			Text:      payload.Text,
			Timestamp: time.Now(),
		}
		ThoughtInputChan <- t // ✅ send to worker

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "✅ Thought saved! thank you")
	})

	// GET /summary
	http.HandleFunc("/summary", func(w http.ResponseWriter, r *http.Request) {
		thoughts := storage.LoadThoughts()
		counts := make(map[string]int)
		for _, t := range thoughts {
			counts[t.Category]++
		}
		result := "📊 Thought Summary:\n"
		for cat, cnt := range counts {
			result += fmt.Sprintf(" - %s: %d\n", cat, cnt)
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(result))
	})

	// GET /summary-detail
	http.HandleFunc("/summary-detail", func(w http.ResponseWriter, r *http.Request) {
		thoughts := storage.LoadThoughts()
		sort.SliceStable(thoughts, func(i, j int) bool {
			return thoughts[i].Timestamp.Before(thoughts[j].Timestamp)
		})

		result := "🧠 All Thoughts:\n"
		for _, t := range thoughts {
			ts := t.Timestamp.Format("Jan 2 15:04")
			result += fmt.Sprintf("• [%s][%s] %s — %s\n", t.Category, t.Tone, t.Text, ts)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(result))
	})

	// GET /prompt
	http.HandleFunc("/prompt", func(w http.ResponseWriter, r *http.Request) {
		p := prompts.GetRandomPrompt()
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(p))
	})

	// GET /export
	http.HandleFunc("/export", func(w http.ResponseWriter, r *http.Request) {
		storage.ExportToFile()
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("💾 Thoughts exported to thoughts_export.json"))
	})

	// GET /download
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		filePath := "thoughts_export.json"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=thoughts_export.json")
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, filePath)
	})

	// 🧠 Auto-update report every 50 seconds using goroutine
	go func() {
		for {
			thoughts := storage.LoadThoughts()
			sort.SliceStable(thoughts, func(i, j int) bool {
				return thoughts[i].Priority > thoughts[j].Priority
			})

			limit := 3
			if len(thoughts) < limit {
				limit = len(thoughts)
			}

			result := "🔥 Top Priority Thoughts:\n"
			for i := 0; i < limit; i++ {
				t := thoughts[i]
				result += fmt.Sprintf(" - [%s] %s (Priority: %d)\n", t.Category, t.Text, t.Priority)
			}
			cachedReport = result
			fmt.Println("⏳ Report refreshed... Next in 50s")
			time.Sleep(50 * time.Second)
		}
	}()

	// GET /report — just return the latest cached report
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(cachedReport))
	})

	// Start server
	fmt.Println("🌐 Web UI running at: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, "❌ Web server failed:", err)
		os.Exit(1)
	}
}
