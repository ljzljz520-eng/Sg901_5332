package api

import (
	"example.com/edu-leaderboard/ranking"
	"example.com/edu-leaderboard/service"
	"example.com/edu-leaderboard/storage"
	"example.com/edu-leaderboard/workflow"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	h := New(workflow.New(service.New(s, p, p))).Routes()
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
