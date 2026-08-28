package ranking

import (
	"example.com/edu-leaderboard/model"
	"strings"
)

func FilterRegion(rs []model.Record, region string, profiles map[string]model.Profile) []model.Record {
	if region == "" {
		return rs
	}
	out := []model.Record{}
	for _, r := range rs {
		if p, ok := profiles[r.PlayerID]; ok && strings.EqualFold(p.Region, region) {
			out = append(out, r)
		}
	}
	return out
}
func Top(rs []model.Record, n int) []model.Record {
	if n < 1 {
		return []model.Record{}
	}
	if n > len(rs) {
		n = len(rs)
	}
	return rs[:n]
}
func Scores(rs []model.Record) []int64 {
	out := make([]int64, len(rs))
	for i, r := range rs {
		out[i] = r.Score
	}
	return out
}
func HasPlayer(rs []model.Record, id string) bool {
	for _, r := range rs {
		if r.PlayerID == id {
			return true
		}
	}
	return false
}
