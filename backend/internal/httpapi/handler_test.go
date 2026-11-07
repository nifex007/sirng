package httpapi

import (
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	for _, method := range []string{"GET", "POST"} {
		w := httptest.NewRecorder()
		NewHandler().ServeHTTP(w, httptest.NewRequest(method, "/api/health", nil))
		if method == "GET" {
			if w.Code != 200 || w.Body.String() != `{"status":"ok"}` || w.Header().Get("Content-Type") != "application/json" {
				t.Fatalf("unexpected response: %v", w)
			}
		} else if w.Code != 405 {
			t.Fatalf("expected 405, got %d", w.Code)
		}
	}
}
