package ranking

import (
	"example.com/edu-leaderboard/model"
	"math"
	"sort"
)

type ScoreBand struct {
	Low, High int64
	Label     string
}

func ScorePercentile(rs []model.Record, score int64) float64 {
	if len(rs) == 0 {
		return 0
	}
	below := 0
	for _, r := range rs {
		if r.Score <= score {
			below++
		}
	}
	return float64(below) / float64(len(rs)) * 100
}
func NormalizeScores(rs []model.Record) []float64 {
	if len(rs) == 0 {
		return []float64{}
	}
	max := rs[0].Score
	for _, r := range rs {
		if r.Score > max {
			max = r.Score
		}
	}
	out := make([]float64, len(rs))
	if max == 0 {
		return out
	}
	for i, r := range rs {
		out[i] = math.Round(float64(r.Score)/float64(max)*1000) / 10
	}
	return out
}
func Bands() []ScoreBand {
	return []ScoreBand{{0, 39, "beginner"}, {40, 69, "developing"}, {70, 89, "advanced"}, {90, 1000000, "expert"}}
}
func BandFor(score int64) ScoreBand {
	for _, b := range Bands() {
		if score >= b.Low && score <= b.High {
			return b
		}
	}
	return Bands()[0]
}
func StableSort(rs []model.Record) []model.Record {
	out := append([]model.Record{}, rs...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out
}
func Position(rs []model.Record, id string) int {
	for i, r := range Rank(rs) {
		if r.ID == id {
			return i + 1
		}
	}
	return 0
}
func Median(rs []model.Record) int64 {
	if len(rs) == 0 {
		return 0
	}
	x := StableSort(rs)
	return x[len(x)/2].Score
}
func Above(rs []model.Record, threshold int64) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.Score >= threshold {
			out = append(out, r)
		}
	}
	return out
}
func Below(rs []model.Record, threshold int64) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.Score < threshold {
			out = append(out, r)
		}
	}
	return out
}
