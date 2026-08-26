package model

import "time"

type Record struct {
	ID, GameID, PlayerID, Source string
	Score                        int64
	Rank                         int
	CreatedAt                    time.Time
	Archived                     bool
}
type Profile struct {
	ID, DisplayName, Region string
	Games                   int
	Active                  bool
}
type Event struct {
	ID, RecordID, Kind, Payload string
	At                          time.Time
}
type Audit struct {
	ID, Action, Actor, Target string
	At                        time.Time
	Success                   bool
}

func NewRecord(id, game, player, source string, score int64) Record {
	return Record{ID: id, GameID: game, PlayerID: player, Source: source, Score: score, CreatedAt: time.Now().UTC()}
}
func (r Record) Valid() bool  { return r.ID != "" && r.GameID != "" && r.PlayerID != "" && r.Score >= 0 }
func (r Record) Key() string  { return r.GameID + ":" + r.PlayerID + ":" + r.ID }
func (p Profile) Valid() bool { return p.ID != "" && p.DisplayName != "" }
func (e Event) Valid() bool   { return e.ID != "" && e.Kind != "" }
func (a Audit) Valid() bool   { return a.ID != "" && a.Action != "" }
