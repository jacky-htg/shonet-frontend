package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
)

func GetUserByID (id uint) (models.User, error) {
	var query string

	query  = "SELECT `id`, `name`, `email`, `password`, `group_id`, `is_active`, `is_reset_password`, `gender`, `phone_number`, `photo`, `birthdate`,`biography`, `city_id`, `created_at`, `updated_at` FROM `users` WHERE `users`.`id` = ?"
	return getUserRow(db.Query(query, id))
}

func getUserRow(rows *sql.Rows, err error) (models.User, error) {

	var user models.User
	libraries.CheckError(err)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var userNull models.UserNull
		err := rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.Group.ID,
				&user.IsActive,
				&user.IsResetPassword,
				&user.Gender,
				&userNull.PhoneNumber,
				&userNull.Photo,
				&userNull.Birthdate,
				&userNull.Biography,
				&userNull.CityId,
				&userNull.CreatedAt,
				&userNull.UpdatedAt,
			)

		if err != nil {
			return models.User{}, err
		}

		user.Password = []byte{}
		user.PhoneNumber = userNull.PhoneNumber.String
		user.Photo = userNull.Photo.String
		user.Biography = userNull.Biography.String
		user.Birthdate = userNull.Birthdate.Time
		user.City.ID = uint(userNull.CityId.Int64)
		user.CreatedAt = userNull.CreatedAt.Time
		user.UpdatedAt = userNull.UpdatedAt.Time
	}

	err = rows.Err()
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}