package web

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func StartServer(port string) error {
	http.HandleFunc("/api/sessions", HandleSessions)
	fmt.Printf("Web server running at http://localhost:%s\n", port)
	fmt.Println("Press Ctrl+C to stop.")
	return http.ListenAndServe(":"+port, nil)
}
