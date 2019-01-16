package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/models"
	"strings"
)

func GetTopBrands(flag int) ([]models.Brand, error) {
	var oriLimit = 5
	var curLimit = 0
	var brandFlag = "fashion_brand"
	var sql string
	var whereNotIn []string

	if flag == 2 {brandFlag = "beauty_brand"}
	sql = "SELECT * FROM `brands` WHERE " +brandFlag+ " = " +string(flag)+ " LIMIT " +string(oriLimit)+ " ORDER BY `brands`.`name` ASC"

	brands, err := fetchBrands(db.Query(sql))
	if err != nil {
		return []models.Brand{}, err
	}

	if len(brands) > 0 {
		curLimit += len(brands)
	}

	if curLimit < oriLimit {
		whereProduct := []string{"`products`.`category_id` IN (SELECT `categories`.`id` FROM `categories` WHERE `categories`.`parent_id` = " +string(flag)+ " )"}

		if len(brands) > 0 {
			for _, val := range brands {whereNotIn = append(whereNotIn, string(val.ID))}
			whereProduct = append(whereProduct, "`products`.`brand_id` NOT IN (" +strings.Join(whereNotIn, ", ")+ ")")
		}

		sql  = " SELECT ViewProduct.brand_id FROM (SELECT products.id, COUNT(products.id) N, products.brand_id FROM log_view_products " +
			   " JOIN products ON (log_view_products.product_id = products.id AND "  +strings.Join(whereProduct," AND ")+ ")" +
			   " GROUP BY products.id " +
			   " ORDER BY N DESC) ViewProduct GROUP BY ViewProduct.brand_id ORDER BY SUM(ViewProduct.N) DESC"

		//------
		//take current brands from current brands
		sql = " SELECT * FROM `brands` WHERE `brands`.`id` NOT IN ( " +strings.Join(whereNotIn, ", ")+ " ) LIMIT " +string(oriLimit - curLimit)+ " ORDER BY `brands`.`name` ASC"
		brands1, err := fetchBrands(db.Query(sql))
		if err != nil {
			return []models.Brand{}, err
		}

		brands = append(brands, brands1...)
	}

	return brands, nil
}

func fetchBrands(rows *sql.Rows, err error) ([]models.Brand, error) {
	var brands []models.Brand

	if err != nil {
		return []models.Brand{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand models.Brand
		var brandNull models.BrandNull

		err = rows.Scan(
				&brand.ID,
				&brand.Name,
				&brandNull.Description,
				&brandNull.Image,
				&brandNull.SiteURL,
				&brandNull.BeautyBrand,
				&brandNull.FashionBrand,
				&brandNull.DeliveryNote,
				&brandNull.ReturnNote,
				&brandNull.SocialMedia,
				&brandNull.VendorTitle,
				&brandNull.CreatedAt,
				&brandNull.UpdatedAt,
			)

		if err != nil {
			return []models.Brand{}, err
		}

		brand.Description = brandNull.Description.String
		brand.Image = brandNull.Image.String
		brand.SiteURL = brandNull.SiteURL.String
		brand.BeautyBrand = uint(brandNull.BeautyBrand.Int64)
		brand.FashionBrand = uint(brandNull.BeautyBrand.Int64)
		brand.DeliveryNote = brandNull.DeliveryNote.String
		brand.SocialMedia = brandNull.SocialMedia.String
		brand.ReturnNote = brandNull.VendorTitle.String
		brand.CreatedAt = brandNull.CreatedAt.Time
		brand.UpdatedAt = brandNull.UpdatedAt.Time

		brands = append(brands, brand)
	}

	if err = rows.Err(); err != nil {
		return []models.Brand{}, err
	}

	return []models.Brand{}, nil
}