package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
	body := rec.Body.String()
	if !strings.Contains(body, "Прометей") {
		t.Fatal("home page does not contain the studio name")
	}
	if strings.Contains(body, "@") {
		t.Fatal("home page contains an email address")
	}
}

func TestNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/no-such", nil)
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status %d", rec.Code)
	}
}
