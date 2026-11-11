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
	if !strings.Contains(rec.Body.String(), "Прометей") {
		t.Fatal("home page does not contain the studio name")
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
