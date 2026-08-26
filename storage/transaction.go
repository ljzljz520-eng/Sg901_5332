package storage

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
)

type Transaction struct {
	s       *Store
	ctx     context.Context
	records []model.Record
	audits  []model.Audit
}

func (s *Store) Begin(ctx context.Context) *Transaction { return &Transaction{s: s, ctx: ctx} }
func (t *Transaction) AddRecord(r model.Record) error {
	if t.ctx.Err() != nil {
		return t.ctx.Err()
	}
	if !r.Valid() {
		return errors.New("invalid record")
	}
	t.records = append(t.records, r)
	return nil
}
func (t *Transaction) AddAudit(a model.Audit) error {
	if !a.Valid() {
		return errors.New("invalid audit")
	}
	t.audits = append(t.audits, a)
	return nil
}
func (t *Transaction) Commit() error {
	for _, r := range t.records {
		if e := t.s.PutRecord(t.ctx, r); e != nil {
			return e
		}
	}
	for _, a := range t.audits {
		if e := t.s.PutAudit(t.ctx, a); e != nil {
			return e
		}
	}
	return nil
}
func (t *Transaction) Rollback() { t.records = nil; t.audits = nil }
