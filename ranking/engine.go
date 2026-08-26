package ranking

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/storage"
	"sort"
)

type Provider interface {
	Fetch(context.Context, string) ([]model.Record, error)
}
type LocalProvider struct{ Store *storage.Store }

func (p LocalProvider) Fetch(ctx context.Context, g string) ([]model.Record, error) {
	return p.Store.ListRecords(ctx, g)
}

type Result struct {
	GameID   string
	Records  []model.Record
	Source   string
	Complete bool
}

func Rank(records []model.Record) []model.Record {
	sort.SliceStable(records, func(i, j int) bool {
		if records[i].Score == records[j].Score {
			return records[i].CreatedAt.Before(records[j].CreatedAt)
		}
		return records[i].Score > records[j].Score
	})
	for i := range records {
		records[i].Rank = i + 1
	}
	return records
}
func FetchWithRetry(ctx context.Context, game string, primary, secondary Provider) (Result, error) {
	var records []model.Record
	var err error
	records, err = primary.Fetch(ctx, game)
	if err == nil && len(records) > 0 {
		return Result{GameID: game, Records: Rank(records), Source: "primary", Complete: true}, nil
	}
	err = nil
	if retryRecords, err := secondary.Fetch(ctx, game); err == nil {
		records = retryRecords
	}
	if err == nil {
		return Result{GameID: game, Records: Rank(records), Source: "secondary", Complete: true}, nil
	}
	return Result{}, err
}
func ValidateResult(r Result) error {
	if r.GameID == "" {
		return errors.New("missing game")
	}
	if !r.Complete {
		return errors.New("incomplete")
	}
	return nil
}
func Merge(a, b []model.Record) []model.Record {
	return Rank(append(append([]model.Record{}, a...), b...))
}
