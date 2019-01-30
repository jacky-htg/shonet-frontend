package models

import (
"database/sql"
"time"

"github.com/go-sql-driver/mysql"
)

type User struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        []byte    `json:"password"`
	Group           Group     `json:"group"`
	IsActive        bool      `json:"is_active"`
	IsResetPassword bool      `json:"is_reset_password"`
	PhoneNumber     string    `json:"phone_number"`
	Photo           string    `json:"photo"`
	Biography       string    `json:"biography"`
	Birthdate       time.Time `json:"birthdate,string,omitempty"`
	Gender          string    `json:"gender"`
	City            City      `json:"city"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Journals        int       `json:"journals"`
	ProductArticles int       `json:"product_articles"`
}

type UserNull struct {
	PhoneNumber sql.NullString
	Photo       sql.NullString
	Biography   sql.NullString
	Birthdate   mysql.NullTime
	CityId      sql.NullInt64
	CreatedAt   mysql.NullTime
	UpdatedAt   mysql.NullTime
	DeletedAt	mysql.NullTime
	CityName	sql.NullString
	CityLat		sql.NullFloat64
	CityLong	sql.NullFloat64
	CountryId	sql.NullInt64
	CountryName	sql.NullString
	CountryCode sql.NullString
}

type UserElasticSearch struct {
	ID 				uint			`json:"id"`
	Name			string			`json:"name"`
	Email			string			`json:"email"`

	Group			struct {
		ID			uint			`json:"id"`
		Title		string			`json:"title"`
	}								`json:"group"`

	IsActive		int32			`json:"is_active"`
	IsResetPassword int32			`json:"is_reset_password"`
	PhoneNumber		string			`json:"phone_number"`
	Photo			string			`json:"photo"`
	Biography 		string			`json:"biography"`
	Birthdate		string			`json:"birthdate"`
	Gender			string			`json:"gender"`

	City			struct {
		ID			uint			`json:"id"`
		Name		string			`json:"name"`
		Country     struct {
			ID		uint			`json:"id"`
			Name    string          `json:"name"`
			Code    string          `json:"code"`
		}							`json:"country"`

		Latitude	float64			`json:"latitude"`
		Longitude	float64			`json:"longitude"`
	}

	CreatedAt		string			`json:"created_at"`
	UpdatedAt       string          `json:"updated_at"`
	DeletedAt       string          `json:"deleted_at"`
	LoginType       string          `json:"login_type"`
	IsDefault		int32			`json:"is_default"`
}

func UserValidation(user User) (User, error) {
	// todo : sanitation and validation of user model
	return user, nil
}

