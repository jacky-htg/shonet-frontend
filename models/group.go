package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type Group struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupNull struct {
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}

