package service

import (
	"context"
	"errors"
	"example.com/edu-leaderboard/model"
	"sort"
	"strings"
)

type Catalog struct {
	games   map[string]string
	regions map[string]bool
}

func NewCatalog() *Catalog { return &Catalog{games: map[string]string{}, regions: map[string]bool{}} }
func (c *Catalog) AddGame(id, name string) error {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(name) == "" {
		return errors.New("game required")
	}
	c.games[id] = name
	return nil
}
func (c *Catalog) RemoveGame(id string) bool {
	if _, ok := c.games[id]; !ok {
		return false
	}
	delete(c.games, id)
	return true
}
func (c *Catalog) GameName(id string) string { return c.games[id] }
func (c *Catalog) Games() []string {
	out := []string{}
	for id := range c.games {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
func (c *Catalog) AddRegion(r string) error {
	r = strings.TrimSpace(r)
	if r == "" {
		return errors.New("region required")
	}
	c.regions[strings.ToLower(r)] = true
	return nil
}
func (c *Catalog) HasRegion(r string) bool { return c.regions[strings.ToLower(strings.TrimSpace(r))] }
func (c *Catalog) ValidateRecord(r model.Record) error {
	if !r.Valid() {
		return errors.New("invalid")
	}
	if _, ok := c.games[r.GameID]; !ok {
		return errors.New("unknown game")
	}
	return nil
}
func (c *Catalog) Seed(ctx context.Context) error {
	if e := ctx.Err(); e != nil {
		return e
	}
	for _, g := range []string{"math", "science", "reading", "coding", "history"} {
		if e := c.AddGame(g, strings.Title(g)); e != nil {
			return e
		}
	}
	return nil
}
func (c *Catalog) Count() int       { return len(c.games) }
func (c *Catalog) RegionCount() int { return len(c.regions) }
func (c *Catalog) Clear()           { c.games = map[string]string{}; c.regions = map[string]bool{} }
