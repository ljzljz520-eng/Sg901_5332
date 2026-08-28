package service

import (
	"context"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/ranking"
	"example.com/edu-leaderboard/storage"
	"path/filepath"
	"testing"
)

func TestRegisterAndQuery(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	x := New(s, p, p)
	if e := x.Register(context.Background(), model.NewRecord("1", "g", "p", "x", 4)); e != nil {
		t.Fatal(e)
	}
	r, e := x.Query(context.Background(), "g")
	if e != nil || len(r.Records) != 1 {
		t.Fatalf("%v", e)
	}
}
