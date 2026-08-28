package main

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/ranking"
	"testing"
)

type emptyProvider struct{}

func (emptyProvider) Fetch(context.Context, string) ([]model.Record, error) {
	return []model.Record{}, nil
}
func TestLeaderboardUnavailableOnRetryFailure(t *testing.T) {
	_, e := ranking.FetchWithRetry(context.Background(), "g", emptyProvider{}, providerErr{})
	if e == nil || !errors.Is(e, errSentinel) {
		t.Fatalf("expected preserved error, got %v", e)
	}
}

var errSentinel = errors.New("source down")

type providerErr struct{}

func (providerErr) Fetch(context.Context, string) ([]model.Record, error) { return nil, errSentinel }
