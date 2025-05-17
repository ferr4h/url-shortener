package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCORSNoOrigin(t *testing.T) {
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusTeapot)
	})

	handler := CORS(next)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if !called {
		t.Error("next handler was not called")
	}

	if rr.Code != http.StatusTeapot {
		t.Errorf("expected status code %d, got %d", http.StatusTeapot, rr.Code)
	}

	if origin := rr.Header().Get("Access-Control-Allow-Origin"); origin != "" {
		t.Errorf("expected no Access-Control-Allow-Origin header, got %q", origin)
	}
}

func TestCORSOptionsRequest(t *testing.T) {
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	handler := CORS(next)

	req := httptest.NewRequest("OPTIONS", "http://example.com/foo", nil)
	req.Header.Set("Origin", "http://origin.example.com")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if called {
		t.Error("next handler should NOT be called on OPTIONS request")
	}

	if got := rr.Header().Get("Access-Control-Allow-Origin"); got != "http://origin.example.com" {
		t.Errorf("expected Access-Control-Allow-Origin header to be set to Origin header, got %q", got)
	}

	if got := rr.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
		t.Errorf("expected Access-Control-Allow-Credentials header to be 'true', got %q", got)
	}

	if got := rr.Header().Get("Access-Control-Allow-Methods"); !strings.Contains(got, "POST") {
		t.Errorf("expected Access-Control-Allow-Methods header to contain 'POST', got %q", got)
	}

	if got := rr.Header().Get("Access-Control-Max-Age"); got != "86400" {
		t.Errorf("expected Access-Control-Max-Age header to be '86400', got %q", got)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestCORSNormalRequestWithOrigin(t *testing.T) {
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusAccepted) // 202
	})

	handler := CORS(next)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Set("Origin", "http://origin.example.com")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if !called {
		t.Error("next handler was not called")
	}

	if got := rr.Header().Get("Access-Control-Allow-Origin"); got != "http://origin.example.com" {
		t.Errorf("expected Access-Control-Allow-Origin header to be set to Origin header, got %q", got)
	}

	if got := rr.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
		t.Errorf("expected Access-Control-Allow-Credentials header to be 'true', got %q", got)
	}

	if rr.Code != http.StatusAccepted {
		t.Errorf("expected status code %d, got %d", http.StatusAccepted, rr.Code)
	}
}
