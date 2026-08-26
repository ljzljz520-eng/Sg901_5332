package service

import "context"

type Metrics struct{ Records, Profiles, Events, Audits int }

func (s *Service) Metrics(ctx context.Context) (Metrics, error) {
	a := Metrics{}
	var e error
	if a.Records, e = s.Store.Count(ctx, "records"); e != nil {
		return a, e
	}
	a.Profiles, _ = s.Store.Count(ctx, "profiles")
	a.Events, _ = s.Store.Count(ctx, "events")
	a.Audits, _ = s.Store.Count(ctx, "audits")
	return a, nil
}
func (m Metrics) Total() int { return m.Records + m.Profiles + m.Events + m.Audits }
