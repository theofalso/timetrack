package session

import (
	"time"
)

type Session struct {
	Project   string    `json:"project"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
}

// IsActive true if session is active (EndTime is zero), false otherwise
func (s *Session) IsActive() bool {
	return s.EndTime.IsZero()
}

// duration returns the duration of the session. If the session is active, it calculates the duration from StartTime to now. If the session has ended, it calculates the duration from StartTime to EndTime.
func (s *Session) Duration() time.Duration {
	if s.IsActive() {
		return time.Since(s.StartTime)
	}
	return s.EndTime.Sub(s.StartTime)
}

// NewSession to create a new session with the given project name.
func NewSession(projectName string) *Session {
	return &Session{
		Project:   projectName,
		StartTime: time.Now(),
	}
}
