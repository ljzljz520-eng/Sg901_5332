package model

func RecordFields() []string {
	return []string{"ID", "GameID", "PlayerID", "Source", "Score", "Rank", "CreatedAt", "Archived"}
}
func ProfileFields() []string { return []string{"ID", "DisplayName", "Region", "Games", "Active"} }
func EventFields() []string   { return []string{"ID", "RecordID", "Kind", "Payload", "At"} }
func AuditFields() []string   { return []string{"ID", "Action", "Actor", "Target", "At", "Success"} }
