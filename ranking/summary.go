package ranking

import (
	"example.com/edu-leaderboard/model"
	"fmt"
)

func Summary(rs []model.Record) string {
	if len(rs) == 0 {
		return "empty"
	}
	return fmt.Sprintf("%d records, top score %d", len(rs), rs[0].Score)
}
func Deduplicate(rs []model.Record) []model.Record {
	seen := map[string]bool{}
	out := []model.Record{}
	for _, r := range rs {
		if !seen[r.ID] {
			seen[r.ID] = true
			out = append(out, r)
		}
	}
	return out
}
