package api

import (
	"encoding/json"
	"example.com/edu-leaderboard/model"
	"net/http"
)

func (s *Server) Stats(w http.ResponseWriter, r *http.Request) {
	m, e := s.Chain.Svc.Metrics(r.Context())
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(m)
}
func (s *Server) Archive(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if e := s.Chain.Archive(r.Context(), id); e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	w.WriteHeader(204)
}
func (s *Server) Event(w http.ResponseWriter, r *http.Request) {
	e := model.Event{ID: r.URL.Query().Get("id"), Kind: r.URL.Query().Get("kind")}
	if err := s.Chain.Track(r.Context(), e); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	w.WriteHeader(201)
}
func (s *Server) AddAdminRoutes(m *http.ServeMux) *http.ServeMux {
	m.HandleFunc("/stats", s.Stats)
	m.HandleFunc("/archive", s.Archive)
	m.HandleFunc("/event", s.Event)
	return m
}
