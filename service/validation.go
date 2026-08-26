package service

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
)

func ValidateBatch(ctx context.Context, rs []model.Record) error {
	if len(rs) == 0 {
		return errors.New("empty batch")
	}
	for _, r := range rs {
		if e := ctx.Err(); e != nil {
			return e
		}
		if !r.Valid() {
			return errors.New("invalid batch record")
		}
	}
	return nil
}
func NormalizeRecord(r model.Record) model.Record {
	if r.Source == "" {
		r.Source = "submission"
	}
	if r.Score < 0 {
		r.Score = 0
	}
	return r
}
func CanArchive(r model.Record) bool         { return !r.Archived && r.ID != "" }
func IsFresh(r model.Record, now int64) bool { return now-r.CreatedAt.Unix() < 86400 }
