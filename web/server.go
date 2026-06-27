package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

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
	http.HandleFunc("/", HandleDashboard)

	fmt.Printf("Web server running at http://localhost:%s\n", port)
	fmt.Println("Press Ctrl+C to stop.")

	return http.ListenAndServe(":"+port, nil)
}
