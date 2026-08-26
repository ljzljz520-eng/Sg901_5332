package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareHeaders(t *testing.T) {
	h := WithJSON(WithRequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
	if w.Header().Get("Content-Type") == "" || w.Header().Get("X-Request-ID") == "" {
		t.Fatal("headers missing")
	}
}
