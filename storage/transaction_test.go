package storage

import (
	"context"
	"example.com/edu-leaderboard/model"
	"path/filepath"
	"testing"
)

func TestTransactionCommit(t *testing.T) {
	s, _ := Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	tx := s.Begin(context.Background())
	if e := tx.AddRecord(model.NewRecord("x", "g", "p", "", 1)); e != nil {
		t.Fatal(e)
	}
	if e := tx.Commit(); e != nil {
		t.Fatal(e)
	}
}
