package ranking

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"testing"
)

type failProvider struct{ err error }

func (f failProvider) Fetch(context.Context, string) ([]model.Record, error) { return nil, f.err }
func TestRetryPropagatesError(t *testing.T) {
	_, e := FetchWithRetry(context.Background(), "g", failProvider{errors.New("primary")}, failProvider{errors.New("secondary")})
	if e == nil {
		t.Fatal("expected error")
	}
}
