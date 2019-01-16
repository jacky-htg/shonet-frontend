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
}

func UserValidation(user User) (User, error) {
	// todo : sanitation and validation of user model
	return user, nil
}

