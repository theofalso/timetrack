package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/theofalso/timetrack/internal/session"
)

// dataPath to save sesssion
const dataPath = "data/sessions.json"

// read JSON and return a slice of Session structs. If the file doesn't exist, it returns an empty slice
func Load() ([]session.Session, error) {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []session.Session{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []session.Session{}, nil
	}

	// to transform the JSON text into Go structures, we use json.Unmarshal
	var sessions []session.Session
	if err := json.Unmarshal(data, &sessions); err != nil {
		return nil, err
	}

	return sessions, nil
}

// Save the slice of Session structs to a JSON file. It overwrites the file if it already exists
func Save(sessions []session.Session) error {
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataPath, data, 0644)
}
