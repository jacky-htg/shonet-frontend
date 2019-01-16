package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Brand struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	SiteURL      string    `json:"site_url"`
	BeautyBrand  uint      `json:"beauty_brand"`
	FashionBrand uint      `json:"fashion_brand"`
	DeliveryNote string    `json:"delivery_note"`
	ReturnNote   string    `json:"return_note"`
	SocialMedia  string    `json:"social_media"`
	VendorTitle  string    `json:"vendor_title"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BrandNull struct {
	Description  sql.NullString
	Image        sql.NullString
	SiteURL      sql.NullString
	BeautyBrand  sql.NullInt64
	FashionBrand sql.NullInt64
	DeliveryNote sql.NullString
	ReturnNote   sql.NullString
	SocialMedia  sql.NullString
	VendorTitle  sql.NullString
	CreatedAt    mysql.NullTime
	UpdatedAt    mysql.NullTime
}

type SocialMedia struct {
	Line      string
	Facebook  string
	Instagram string
	WhatsApp  string
	Twitter   string
}

func (b *Brand) CheckNull() bool {

	return b.ID == 0
}
