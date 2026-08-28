package main

import (
	"example.com/edu-leaderboard/api"
	"example.com/edu-leaderboard/ranking"
	"example.com/edu-leaderboard/service"
	"example.com/edu-leaderboard/storage"
	"example.com/edu-leaderboard/workflow"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("LEADERBOARD_DB")
	if path == "" {
		path = "leaderboard.db"
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	p := ranking.LocalProvider{Store: s}
	svc := service.New(s, p, p)
	chain := workflow.New(svc)
	log.Fatal(http.ListenAndServe(":8080", api.New(chain).Routes()))
}
