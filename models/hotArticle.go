package models

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type HotArticle struct {
	ID			uint 		`json:"id"`
	Article 	Article		`json:"articles"`
	StartDate	time.Time	`json:"start_date"`
	EndDate		time.Time	`json:"end_date"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type HotArticleNull struct {
	CreatedAt  	mysql.NullTime
	UpdatedAt	mysql.NullTime
}

func (h *HotArticle) CheckNull() bool {
	return h.ID == 0
}

