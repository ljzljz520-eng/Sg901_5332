package storage

import (
	"context"
	"example.com/edu-leaderboard/model"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord("r1", "g", "p", "x", 9)
	if e = s.PutRecord(context.Background(), r); e != nil {
		t.Fatal(e)
	}
	s, e = s.Reopen()
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord(context.Background(), "r1"); e != nil {
		t.Fatal(e)
	}
}
