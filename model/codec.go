package model

import "encoding/json"

func Encode(v any) ([]byte, error)    { return json.Marshal(v) }
func Decode(data []byte, v any) error { return json.Unmarshal(data, v) }
func CloneRecord(r Record) Record     { return r }
func CloneProfile(p Profile) Profile  { return p }
func CloneEvent(e Event) Event        { return e }
func CloneAudit(a Audit) Audit        { return a }
