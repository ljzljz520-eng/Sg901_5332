package storage

import (
	"context"
	"example.com/edu-leaderboard/model"
	"sort"
)

func (s *Store) Recent(ctx context.Context, game string, limit int) ([]model.Record, error) {
	rs, e := s.ListRecords(ctx, game)
	if e != nil {
		return nil, e
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i].CreatedAt.After(rs[j].CreatedAt) })
	if limit > 0 && len(rs) > limit {
		rs = rs[:limit]
	}
	return rs, nil
}
func (s *Store) Archived(ctx context.Context, game string) ([]model.Record, error) {
	rs, e := s.ListRecords(ctx, game)
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.Archived {
			out = append(out, r)
		}
	}
	return out, nil
}
