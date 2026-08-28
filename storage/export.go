package storage

import (
	"context"
	"encoding/json"
	"errors"
	"example.com/edu-leaderboard/model"
	"go.etcd.io/bbolt"
)

func (s *Store) Export(ctx context.Context, game string) ([]byte, error) {
	rs, e := s.ListRecords(ctx, game)
	if e != nil {
		return nil, e
	}
	return json.Marshal(rs)
}
func (s *Store) Import(ctx context.Context, data []byte) error {
	var rs []model.Record
	if e := json.Unmarshal(data, &rs); e != nil {
		return e
	}
	for _, r := range rs {
		if !r.Valid() {
			return errors.New("invalid import")
		}
		if e := s.PutRecord(ctx, r); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) DeleteRecord(ctx context.Context, id string) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
func (s *Store) Exists(ctx context.Context, id string) bool {
	_, e := s.GetRecord(ctx, id)
	return e == nil
}
func (s *Store) Snapshot(ctx context.Context) (map[string]int, error) {
	out := map[string]int{}
	for _, b := range []string{"records", "profiles", "events", "audits"} {
		n, e := s.Count(ctx, b)
		if e != nil {
			return nil, e
		}
		out[b] = n
	}
	return out, nil
}
