package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/models"
	"strings"
)

func GetTopProduct(flag uint) (TopProductMenu, error) {
	var result TopProductMenu
	var oriLimit = 3
	var curLimit = 0
	var whereNotIn = []string{}

	sqlWords := " SELECT * FROM `products` WHERE `products`.`top_product` != 0 AND `products.category_id` IN ( " +
		   		" SELECT * FROM `categories` WHERE `categories`.`parent_id` = " +string(flag)+
		   		" ) ORDER BY `products`.`name` DESC LIMIT " +string(oriLimit)

	products, err := fetchProducts(db.Query(sqlWords))
	if err != nil {
		return TopProductMenu{}, err
	}

	if len(products) > 0 {
		result.Top = products
		for _, val := range products {whereNotIn = append(whereNotIn, string(val.ID))}
		curLimit += len(products)
	}

	if curLimit < oriLimit {
		sqlWords = " SELECT * FROM `products` WHERE `products`.`id` NOT IN ( " +
				   " " +strings.Join(whereNotIn, ", ")+ " ) " +
				   " AND `products`.`category_id` IN ( " +
				   " SELECT `id` FROM `categories` WHERE `categories`.`parent_id` = " +string(flag)+
				   " ) ORDER BY `products`.`view`, `products`.`id` DESC  LIMIT " +string(oriLimit - curLimit)

		products, err = fetchProducts(db.Query(sqlWords))
		if err != nil {
			return TopProductMenu{}, err
		}

		result.Products = products
	}

	return result, nil
}

func fetchProducts(rows *sql.Rows, err error) ([]models.Product, error) {
	var products []models.Product

	if err != nil {
		return []models.Product{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var productNull models.ProductNull

		err = rows.Scan(
				&product.ID,
				&product.Name,
				&product.Brand.ID,
				&product.Category.ID,
				&product.Thumbnail,
				&productNull.Description,
				&product.Price,
				&product.SiteURL,
				&product.MustHave,
				&product.TopProduct,
				&productNull.CreatedAt,
				&productNull.UpdatedAt,
				&productNull.View,
				&productNull.Click,
			)

		if err != nil {
			return []models.Product{}, err
		}

		product.Description = productNull.Description.String
		product.CreatedAt = productNull.CreatedAt.Time
		product.UpdatedAt = productNull.UpdatedAt.Time
		product.View = int(productNull.View.Int64)
		product.Click = int(productNull.Click.Int64)

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return []models.Product{}, err
	}

	return products, nil
}