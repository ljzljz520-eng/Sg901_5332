package service

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/ranking"
	"example.com/edu-leaderboard/storage"
	"fmt"
	"time"
)

type Service struct {
	Store              *storage.Store
	Primary, Secondary ranking.Provider
}

func New(s *storage.Store, p, q ranking.Provider) *Service {
	return &Service{Store: s, Primary: p, Secondary: q}
}
func (s *Service) Register(ctx context.Context, r model.Record) error {
	if !r.Valid() {
		return errors.New("invalid record")
	}
	if r.CreatedAt.IsZero() {
		r.CreatedAt = time.Now().UTC()
	}
	return s.Store.PutRecord(ctx, r)
}
func (s *Service) CreateProfile(ctx context.Context, p model.Profile) error {
	if !p.Valid() {
		return errors.New("invalid profile")
	}
	return s.Store.PutProfile(ctx, p)
}
func (s *Service) RecordEvent(ctx context.Context, e model.Event) error {
	if !e.Valid() {
		return errors.New("invalid event")
	}
	return s.Store.PutEvent(ctx, e)
}
func (s *Service) Query(ctx context.Context, game string) (ranking.Result, error) {
	if game == "" {
		return ranking.Result{}, errors.New("game required")
	}
	r, e := ranking.FetchWithRetry(ctx, game, s.Primary, s.Secondary)
	if e != nil {
		return ranking.Result{}, fmt.Errorf("leaderboard unavailable: %w", e)
	}
	return r, nil
}
func (s *Service) Archive(ctx context.Context, id string) error {
	r, e := s.Store.GetRecord(ctx, id)
	if e != nil {
		return e
	}
	r.Archived = true
	return s.Store.PutRecord(ctx, r)
}
func (s *Service) Audit(ctx context.Context, a model.Audit) error { return s.Store.PutAudit(ctx, a) }
