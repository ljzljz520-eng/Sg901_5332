package workflow

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/service"
)

type Chain struct{ Svc *service.Service }

func New(s *service.Service) *Chain { return &Chain{Svc: s} }
func (c *Chain) Receive(ctx context.Context, r model.Record) error {
	r = service.NormalizeRecord(r)
	if e := service.ValidateBatch(ctx, []model.Record{r}); e != nil {
		return e
	}
	return c.Svc.Register(ctx, r)
}
func (c *Chain) Process(ctx context.Context, game string) (int, error) {
	r, e := c.Svc.Query(ctx, game)
	if e != nil {
		return 0, e
	}
	return len(r.Records), nil
}
func (c *Chain) Archive(ctx context.Context, id string) error   { return c.Svc.Archive(ctx, id) }
func (c *Chain) Track(ctx context.Context, e model.Event) error { return c.Svc.RecordEvent(ctx, e) }
func EnsureChain(c *Chain) error {
	if c == nil || c.Svc == nil {
		return errors.New("chain unavailable")
	}
	return nil
}
