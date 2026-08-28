package workflow

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/service"
)

type BatchResult struct {
	Accepted, Rejected int
	Errors             []string
}

func (c *Chain) SubmitBatch(ctx context.Context, rs []model.Record) BatchResult {
	out := BatchResult{}
	for _, r := range rs {
		if e := c.Receive(ctx, r); e != nil {
			out.Rejected++
			out.Errors = append(out.Errors, e.Error())
		} else {
			out.Accepted++
		}
	}
	return out
}
func (c *Chain) ArchiveBatch(ctx context.Context, ids []string) BatchResult {
	out := BatchResult{}
	for _, id := range ids {
		if e := c.Archive(ctx, id); e != nil {
			out.Rejected++
			out.Errors = append(out.Errors, e.Error())
		} else {
			out.Accepted++
		}
	}
	return out
}
func (c *Chain) Validate(ctx context.Context, r model.Record) error {
	if c == nil || c.Svc == nil {
		return errors.New("service unavailable")
	}
	return service.ValidateBatch(ctx, []model.Record{r})
}
func (c *Chain) ProcessAndArchive(ctx context.Context, game string) error {
	rs, e := c.Svc.Store.ListRecords(ctx, game)
	if e != nil {
		return e
	}
	for _, r := range rs {
		if !r.Archived {
			if e := c.Archive(ctx, r.ID); e != nil {
				return e
			}
		}
	}
	return nil
}
