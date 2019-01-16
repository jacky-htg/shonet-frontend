package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/models"
)

func GetTopProduct() ([]models.Product, error) {
	var result TopProductMenu
	var oriLimit = 3
	var curLimit = 0
	var whereNotIn = []string{}

	sql := ""
	return []models.Product{}, nil
}

func fetchProducts(rows *sql.Rows, err error) ([]models.Product, error) {

	return []models.Product{}, nil
}