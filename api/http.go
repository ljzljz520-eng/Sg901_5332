package api

import (
	"context"
	"encoding/json"
	"example.com/edu-leaderboard/model"
	"example.com/edu-leaderboard/workflow"
	"net/http"
)

type Server struct{ Chain *workflow.Chain }

func New(c *workflow.Chain) *Server { return &Server{Chain: c} }
func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (s *Server) Submit(w http.ResponseWriter, r *http.Request) {
	var rec model.Record
	if json.NewDecoder(r.Body).Decode(&rec) != nil {
		http.Error(w, "bad request", 400)
		return
	}
	if e := s.Chain.Receive(r.Context(), rec); e != nil {
		http.Error(w, e.Error(), 422)
		return
	}
	w.WriteHeader(201)
}
func (s *Server) Leaderboard(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Query().Get("game")
	res, e := s.Chain.Svc.Query(context.Background(), game)
	if e != nil {
		http.Error(w, e.Error(), 503)
		return
	}
	json.NewEncoder(w).Encode(res)
}
func (s *Server) Routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.Health)
	m.HandleFunc("/submit", s.Submit)
	m.HandleFunc("/leaderboard", s.Leaderboard)
	return m
}
