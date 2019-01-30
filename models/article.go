package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Article struct {
	ID                  uint       `json:"id"`
	Title               string     `json:"title"`
	Slug                string     `json:"slug"`
	Permalink           string     `json:"permalink"`
	Content             string     `json:"content"`
	Image               string     `json:"image"`
	ImageSource         string     `json:"image_source"`
	SeoKeyword          string     `json:"seo_keyword"`
	Type                string     `json:"type"`
	Status              string     `json:"status"`
	RequestPublishDate  time.Time  `json:"request_publish_date,omitempty"`
	PublishDate         time.Time  `json:"publish_date,omitempty"`
	Writer              User       `json:"writer"`
	Editor              User       `json:"editor"`
	Tags                []Tag      `json:"tags"`
	Products            []Product  `json:"products"`
	Category            []Category `json:"category"`
	TrendingCount       uint       `json:"trending_count"`
	ContentManipulation string     `json:"content_manipulation"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdateAt            time.Time  `json:"updated_at"`
}

type ArticleNull struct {
	Content             sql.NullString
	Image               sql.NullString
	ImageSource         sql.NullString
	SeoKeyword          sql.NullString
	RequestPublishDate  mysql.NullTime
	PublishDate         mysql.NullTime
	TrendingCount       sql.NullInt64
	CreatedAt           mysql.NullTime
	UpdatedAt           mysql.NullTime
	ContentManipulation sql.NullString
}

type ArticleElastic struct {
	ID					uint			`json:"id"`
	Title				string			`json:"title"`
	Slug                string      	`json:"slug"`
	Permalink           string      	`json:"permalink"`
	Content             string      	`json:"content"`
	Image               string      	`json:"image"`
	ImageSource         string      	`json:"image_source"`
	SeoKeyword          string      	`json:"seo_keyword"`
	Type                string      	`json:"type"`
	Status              string      	`json:"status"`
	RequestPublishDate  string   		`json:"request_publish_date,omitempty"`
	PublishDate         string   		`json:"publish_date,omitempty"`
	Writer              ArticleWriter   `json:"writer"`
	Editor              ArticleWriter   `json:"editor"`
	CreatedAt           string   		`json:"created_at"`
	ContentManipulation string      	`json:"content_manipulation"`
	Tags				[]Tags			`json:"tags"`
	Categories			[]Categories	`json:"categories"`
	Products			[]Products		`json:"products"`
}

type Tags struct {
	ID		uint		`json:"id"`
	Title	string		`json:"title"`
}

type ArticleWriter struct {
	ID 		uint		`json:"id"`
	Name	string		`json:"name"`
	Photo	string		`json:"photo"`
}

type Categories struct {
	ID		uint	`json:"id"`
	Title	string	`json:"title"`
}

type Products struct {
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
	Thumbnail	string	`json:"thumbnail"`
	Price		float64	`json:"price"`
	SiteURL		string	`json:"site_url"`

	Brand	struct {
		ID 		uint	`json:"id"`
		Name	string	`json:"name"`

	} 					`json:"brand"`
}



func (ae *ArticleElastic) IsNull() bool {

	return ae.ID == 0
}

func (a *Article) CheckNull() bool {

	return a.ID == 0
}