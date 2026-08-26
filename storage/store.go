package storage

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"go.etcd.io/bbolt"
	"path/filepath"
	"time"
)

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits")}

type Store struct {
	db   *bbolt.DB
	path string
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(filepath.Clean(path), 0600, &bbolt.Options{Timeout: time.Second})
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func (s *Store) Reopen() (*Store, error) {
	if e := s.Close(); e != nil {
		return nil, e
	}
	return Open(s.path)
}
func (s *Store) PutRecord(ctx context.Context, r model.Record) error {
	return s.put(ctx, "records", r.ID, r)
}
func (s *Store) PutProfile(ctx context.Context, p model.Profile) error {
	return s.put(ctx, "profiles", p.ID, p)
}
func (s *Store) PutEvent(ctx context.Context, e model.Event) error {
	return s.put(ctx, "events", e.ID, e)
}
func (s *Store) PutAudit(ctx context.Context, a model.Audit) error {
	return s.put(ctx, "audits", a.ID, a)
}
func (s *Store) put(ctx context.Context, b, key string, v any) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	data, e := model.Encode(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(b)).Put([]byte(key), data) })
}
func (s *Store) GetRecord(ctx context.Context, id string) (model.Record, error) {
	var v model.Record
	e := s.get(ctx, "records", id, &v)
	return v, e
}
func (s *Store) GetProfile(ctx context.Context, id string) (model.Profile, error) {
	var v model.Profile
	e := s.get(ctx, "profiles", id, &v)
	return v, e
}
func (s *Store) GetEvent(ctx context.Context, id string) (model.Event, error) {
	var v model.Event
	e := s.get(ctx, "events", id, &v)
	return v, e
}
func (s *Store) GetAudit(ctx context.Context, id string) (model.Audit, error) {
	var v model.Audit
	e := s.get(ctx, "audits", id, &v)
	return v, e
}
func (s *Store) get(ctx context.Context, b, id string, v any) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		raw := tx.Bucket([]byte(b)).Get([]byte(id))
		if raw == nil {
			return errors.New("not found")
		}
		return model.Decode(raw, v)
	})
}
func (s *Store) ListRecords(ctx context.Context, game string) ([]model.Record, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if model.Decode(v, &r) != nil {
				return errors.New("decode")
			}
			if game == "" || r.GameID == game {
				out = append(out, r)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) Count(ctx context.Context, b string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		if tx.Bucket([]byte(b)) == nil {
			return errors.New("bucket")
		}
		n = tx.Bucket([]byte(b)).Stats().KeyN
		return nil
	})
	return n, e
}
