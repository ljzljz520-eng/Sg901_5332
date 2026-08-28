package workflow

import (
	"context"
	"example.com/edu-leaderboard/ranking"
)

func (c *Chain) Report(ctx context.Context, game string) (string, error) {
	r, e := c.Svc.Query(ctx, game)
	if e != nil {
		return "", e
	}
	return ranking.Summary(r.Records), nil
}
func (c *Chain) Top(ctx context.Context, game string, n int) ([]int64, error) {
	r, e := c.Svc.Query(ctx, game)
	if e != nil {
		return nil, e
	}
	return ranking.Scores(ranking.Top(r.Records, n)), nil
}
