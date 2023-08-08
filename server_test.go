package fiskil

import (
	"os"
	"testing"
)

func TestServerPortIsConfigured(t *testing.T) {
	want := "80"
	os.Setenv("PORT", want)
	s := Server{}
	s.configurePort()

	if s.Port != want {
		t.Errorf("port not configured, want %v, got %v", want, s.Port)
	} else {
		t.Logf("server port set to %v", want)
	}
	os.Unsetenv("PORT")
}

func TestServerDefaultPort(t *testing.T) {
	want := "8080"
	s := Server{}
	s.configurePort()

	if s.Port != want {
		t.Errorf("port not configured, want %v, got %v", want, s.Port)
	} else {
		t.Logf("server port set to %v", want)
	}
}
