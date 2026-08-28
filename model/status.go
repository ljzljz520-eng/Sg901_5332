package model

type Status string

const (
	StatusPending   Status = "pending"
	StatusProcessed Status = "processed"
	StatusArchived  Status = "archived"
)

func (s Status) Valid() bool {
	return s == StatusPending || s == StatusProcessed || s == StatusArchived
}
func NextStatus(s Status) Status {
	switch s {
	case StatusPending:
		return StatusProcessed
	case StatusProcessed:
		return StatusArchived
	default:
		return s
	}
}
