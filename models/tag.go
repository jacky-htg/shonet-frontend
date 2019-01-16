package models

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type Tag struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TagNull struct {
	CreatedAt mysql.NullTime
	UpdatedAt mysql.NullTime
}