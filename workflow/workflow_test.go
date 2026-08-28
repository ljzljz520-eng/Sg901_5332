package workflow

import (
	"context"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/ranking"
	"example.com/edu-leaderboard/service"
	"example.com/edu-leaderboard/storage"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	c := New(service.New(s, p, p))
	if e := c.Receive(context.Background(), model.NewRecord("1", "g", "p", "", 3)); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	c := New(service.New(s, p, p))
	c.Receive(context.Background(), model.NewRecord("1", "g", "p", "", 3))
	if _, e := c.Process(context.Background(), "g"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	c := New(service.New(s, p, p))
	if e := c.Track(context.Background(), model.Event{ID: "e", Kind: "submit"}); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain17(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	c := New(service.New(s, p, p))
	if e := c.Receive(context.Background(), model.NewRecord("1", "g", "p", "", 3)); e != nil {
		t.Fatal(e)
	}
	if e := c.Archive(context.Background(), "1"); e != nil {
		t.Fatal(e)
	}
}
