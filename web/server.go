package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/theofalso/timetrack/internal/store"
)

func HandleSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sessions, err := store.Load()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not load sessions"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessions)
}

func HandleReportAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	timeRange := r.URL.Query().Get("range")

	sessions, err := store.Load()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not load sessions"})
		return
	}

	var limit time.Time
	now := time.Now()

	switch timeRange {
	case "day":
		limit = now.Add(-24 * time.Hour)
	case "week":
		limit = now.Add(-7 * 24 * time.Hour)
	default:
		limit = time.Time{}
	}

	totals := make(map[string]int64)
	for _, s := range sessions {
		if !limit.IsZero() && s.StartTime.Before(limit) {
			continue
		}
		totals[s.Project] += int64(s.Duration().Seconds())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(totals)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	sessions, err := store.Load()
	if err != nil {
		http.Error(w, "Error loading sessions", http.StatusInternalServerError)
		return
	}

	tmplPath := filepath.Join("web", "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, sessions)
}

func StartServer(port string) error {
	http.HandleFunc("/api/sessions", HandleSessions)
	http.HandleFunc("/api/report", HandleReportAPI)
	http.HandleFunc("/", HandleDashboard)
	fmt.Printf("Web server running securely at http://127.0.0.1:%s\n", port)
	fmt.Println("Press Ctrl+C to stop.")
	return http.ListenAndServe("127.0.0.1:"+port, nil)
}
