package session

import (
	"testing"
	"time"
)

func TestNewSession(t *testing.T) {
	s := NewSession("backend-api")

	if s.Project != "backend-api" {
		t.Errorf("Esperaba el proyecto 'backend-api', pero obtuve '%s'", s.Project)
	}

	if !s.IsActive() {
		t.Errorf("Una sesión recién creada debería estar activa")
	}
}

func TestDuration(t *testing.T) {
	s := NewSession("test")

	s.StartTime = time.Now().Add(-2 * time.Hour)

	s.EndTime = time.Now()

	duracion := s.Duration()

	if duracion.Round(time.Second) != 2*time.Hour {
		t.Errorf("Esperaba una duración de 2h, pero obtuve %v", duracion)
	}
}
