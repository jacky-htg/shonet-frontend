package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Brand       Brand     `json:"brand"`
	Category    Category  `json:"category"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SiteURL     string    `json:"site_url"`
	MustHave    uint      `json:"must_have"`
	TopProduct  uint      `json:"top_product"`
	CreatedBy   int64     `json:"created_by"`
	View        int       `json:"view"`
	Click       int       `json:"click"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductNull struct {
	Description sql.NullString
	CreatedAt   mysql.NullTime
	UpdatedAt   mysql.NullTime
	View		sql.NullInt64
	Click		sql.NullInt64
}

type DescriptionProduct struct {
	EditorsNote   string `json:"editors_note"`
	ProductDetail string `json:"product_detail"`
}
