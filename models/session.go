package models

import (
	"database/sql"
)

type Session struct {
	ID           string `json:"id"`
	IpAddress    string `json:"ip_address"`
	LastActivity int    `json:"last_activity"`
	Payload      string `json:"payload"`
	UserAgent    string `json:"user_agent"`
	UserId       int    `json:"user_id"`
}

type SessionNull struct {
	IpAddress sql.NullString
	UserAgent sql.NullString
	UserId    sql.NullInt64
}
