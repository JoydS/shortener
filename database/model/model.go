package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

const NullStr = "null"

// NullTime to be used in place of sql.NullTime.
type NullTime struct {
	sql.NullTime
}

func NullTimeFromTime(time time.Time) NullTime {
	return NullTime{sql.NullTime{Time: time, Valid: true}}
}

func (s *NullTime) MarshalJSON() ([]byte, error) {
	if s.Valid && !s.Time.IsZero() {
		return json.Marshal(s.Time)
	}

	return []byte(`null`), nil
}

func (s *NullTime) UnmarshalJSON(b []byte) error {
	s.Valid = string(b) != NullStr
	e := json.Unmarshal(b, &s.Time)

	return e
}

// NewNullTime create a new null string. Empty string evaluates to an.
func NewNullTime(value string) *NullTime {
	var null NullTime
	if tval, err := time.Parse(time.RFC3339, value); err == nil {
		null.Time = tval
		null.Valid = true

		return &null
	}

	null.Valid = false

	return &null
}
