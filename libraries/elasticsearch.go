package libraries

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Indexing struct {
	Index struct {
		Index	string		`json:"_index"`
		Type	string		`json:"_type"`
		ID		string		`json:"_id"`
	} `json:"index"`
}

func BulkingAllDataFromSQL(tables map[string]string) (bool, error) {

	db, err := sql.Open(config.GetString("database.mysql.driverName"), config.GetString("database.mysql.dataSourceName"))
	if err != nil {
		return false, err
	}

	err = db.Ping()
	if err != nil {
		return false, err
	}

	defer db.Close()

	counter, err := getCountData(db, tables)
	if err != nil {
		return false, err
	}

	if counter > 0 {
		page := 1
		limit := 100
		nullStatus := false

		for !nullStatus {

			switch tables["name"] {
			case "articles":
				{
					offset := page * limit - limit
					tables["query"] = "  SELECT `articles`.`id`, `articles`.`title`, `articles`.`slug`, `articles`.`permalink`, "+
									  " `articles`.`content`, `articles`.`image`, `articles`.`image_source`, `articles`.`seo_keyword`, "+
									  " `articles`.`type`, `articles`.`status`, `articles`.`request_publish_date`, `articles`.`publish_date`, "+
									  " `articles`.`writer`, `articles`.`editor`, `articles`.`created_at`, `articles`.`content_manipulation`, "+
									  " `writer`.`name`, `writer`.`photo`, `editor`.`name`, `editor`.`photo` "+
									  "  FROM `articles` "+
									  "  LEFT JOIN `users` AS `writer` ON `writer`.`id` = `articles`.`writer` "+
									  "  LEFT JOIN `users` AS `editor` ON `editor`.`id` = `articles`.`editor` "+
									  "  WHERE `articles`.`status` = 'P' AND (`articles`.`publish_date` IS NOT NULL AND `articles`.`publish_date` <= NOW()) "+
									  "  ORDER BY `articles`.`id` DESC LIMIT " +strconv.Itoa(limit)+ " OFFSET " +strconv.Itoa(offset)

					result, err := insertArticlesBulking(db, tables)
					page += 1
					if !result {
						if err != nil {
							nullStatus = true; return false, err
						}

						nullStatus = true
					}
				}
			case "products":
				{

				}
			case "users":
				{
					offset := page * limit - limit
					tables["query"] = "  SELECT `users`.`id`, `users`.`name`, `users`.`email`, "+
									  " `users`.`group_id`, `groups`.`title`, "+
									  " `users`.`is_active`, `users`.`is_reset_password`, `users`.`phone_number`, `users`.`photo`, `users`.`biography`, `users`.`birthdate`, `users`.`gender`, "+
									  " `users`.`city_id`, `city`.`name`, `city`.`latitude`, `city`.`longitude`, "+
									  " `city`.`country_id`, `country`.`name`, `country`.`code`, "+
									  " `users`.`created_at`, `users`.`updated_at`, `users`.`deleted_at`, `users`.`login_type`, `users`.`is_default` "+
									  "  FROM `users` "+
									  "  LEFT JOIN `groups` ON `groups`.`id` = `users`.`group_id` "+
									  "  LEFT JOIN `cities` AS `city` ON `users`.`city_id` = `city`.`id` "+
									  "  LEFT JOIN `countries` AS `country` ON `country`.`id` = `city`.`country_id` "+
									  "  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT " +strconv.Itoa(limit)+ " OFFSET " +strconv.Itoa(offset)

					result, err := insertUsersBulking(db, tables)
					page += 1
					if !result {
						if err != nil {
							nullStatus = true; return false, err
						}

						nullStatus = true
					}
				}
			default:
				{
					err = fmt.Errorf("invalid table name for bulking ::: for information type help")
					return false, err
				}
			}
		}
	}

	return true, nil
}

func getCountData(db *sql.DB, tables map[string]string) (int, error) {

	var sqlString = "SELECT COUNT(`id`) FROM " +tables["name"]
	switch tables["name"] {
	case "articles":
		{
			sqlString += " WHERE `articles`.`status`= 'P' AND (`articles`.`publish_date` IS NOT NULL AND `articles`.`publish_date` <= NOW())"
		}
	case "users":
		{
			sqlString += " WHERE `users`.`deleted_at` IS NULL"
		}
	}

	var counter uint

	row, err := db.Query(sqlString)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	for row.Next() {
		err = row.Scan(&counter)
		if err != nil {
			return 0, err
		}
	}

	err = row.Err()
	if err != nil {
		return 0, err
	}

	return int(counter), nil
}

func insertUsersBulking(db *sql.DB, table map[string]string) (bool, error) {
	var usersList string
	var tglFormat = "2006-01-02 15:04:05"

	rows, err := db.Query(table["query"])
	if err != nil {return false, err}

	if !rows.Next() {return false, nil}

	for rows.Next() {
		var userElastic     models.UserElasticSearch
		var usersNull       models.UserNull
		var index			Indexing

		err = rows.Scan(
				&userElastic.ID,
				&userElastic.Name,
				&userElastic.Email,
				&userElastic.Group.ID,
				&userElastic.Group.Title,
				&userElastic.IsActive,
				&userElastic.IsResetPassword,
				&usersNull.PhoneNumber,
				&usersNull.Photo,
				&usersNull.Biography,
				&usersNull.Birthdate,
				&userElastic.Gender,
				&usersNull.CityId,
				&usersNull.CityName,
				&usersNull.CityLat,
				&usersNull.CityLong,
				&usersNull.CountryId,
				&usersNull.CountryName,
				&usersNull.CountryCode,
				&usersNull.CreatedAt,
				&usersNull.UpdatedAt,
				&usersNull.DeletedAt,
				&userElastic.LoginType,
				&userElastic.IsDefault,
			)

		if err != nil {return false, err}

		userElastic.PhoneNumber 		= usersNull.PhoneNumber.String
		userElastic.Photo 				= usersNull.Photo.String
		userElastic.Biography 			= usersNull.Biography.String
		userElastic.Birthdate			= usersNull.Birthdate.Time.Format(tglFormat)
		userElastic.City.ID				= uint(usersNull.CityId.Int64)
		userElastic.City.Name			= usersNull.CityName.String
		userElastic.City.Latitude   	= usersNull.CityLat.Float64
		userElastic.City.Longitude  	= usersNull.CityLong.Float64
		userElastic.City.Country.ID 	= uint(usersNull.CountryId.Int64)
		userElastic.City.Country.Name	= usersNull.CountryName.String
		userElastic.City.Country.Code   = usersNull.CountryCode.String
		userElastic.CreatedAt			= usersNull.CreatedAt.Time.Format(tglFormat)
		userElastic.UpdatedAt			= usersNull.UpdatedAt.Time.Format(tglFormat)
		userElastic.DeletedAt			= usersNull.DeletedAt.Time.Format(tglFormat)

		index.Index.Index = config.GetString("database.elasticsearch.prefix") + table["name"]
		index.Index.Type  = table["name"]
		index.Index.ID    = strconv.Itoa(int(userElastic.ID))

		indexing, err := json.Marshal(index)
		if err != nil {return false, err}

		users, err := json.Marshal(userElastic)
		if err != nil {return false, err}

		usersList += string(indexing) + "\n"
		usersList += string(users)    + "\n"
	}

	if err = rows.Err(); err != nil {return false, err}

	result, err := insertDataBulking(usersList, table["name"])
	if err != nil {return result, err}

	return true, nil
}

func insertArticlesBulking(db *sql.DB, tables map[string]string) (bool, error) {
	var articleLists string
	var tglFormat = "2006-01-02 15:04:05"

	rows, err := db.Query(tables["query"])
	if err != nil {return false, err}

	defer rows.Close()
	if !rows.Next() {return false, nil}

	for rows.Next() {
		var articleElastics models.ArticleElastic
		var article 		models.Article
		var indexed 		Indexing

		err = rows.Scan(
				&articleElastics.ID,
				&articleElastics.Title,
				&articleElastics.Slug,
				&articleElastics.Permalink,
				&articleElastics.Content,
				&articleElastics.Image,
				&articleElastics.ImageSource,
				&articleElastics.SeoKeyword,
				&articleElastics.Type,
				&articleElastics.Status,
				&article.RequestPublishDate,
				&article.PublishDate,
				&articleElastics.Writer.ID,
				&articleElastics.Editor.ID,
				&article.CreatedAt,
				&articleElastics.ContentManipulation,
				&articleElastics.Writer.Name,
				&articleElastics.Writer.Photo,
				&articleElastics.Editor.Name,
				&articleElastics.Editor.Photo,
			)

		articleElastics.RequestPublishDate 	= article.RequestPublishDate.Format(tglFormat)
		articleElastics.PublishDate 		= article.PublishDate.Format(tglFormat)
		articleElastics.CreatedAt 			= article.CreatedAt.Format(tglFormat)

		if articleElastics.Editor.ID == 0 {
			articleElastics.Editor.ID	 	= articleElastics.Writer.ID
			articleElastics.Editor.Name 	= articleElastics.Writer.Name
			articleElastics.Editor.Photo	= articleElastics.Writer.Photo
		}

		if articleElastics.ContentManipulation == "" {
			articleElastics.ContentManipulation, err = MediaManipulation(articleElastics.Content)
			if err != nil {return false, err}
		}

		if err != nil {return false, err}

		tags, err := getTagsArticle(db, articleElastics.ID)
		if err != nil {return false, err}
		articleElastics.Tags = tags

		categories, err := getCategoriesArticle(db, articleElastics.ID)
		if err != nil {return false, err}
		articleElastics.Categories = categories

		products, err := getProductsArticle(db, articleElastics.ID)
		if err != nil {return false, err}
		articleElastics.Products = products

		indexed.Index.Index 	= config.GetString("database.elasticsearch.prefix") + tables["name"]
		indexed.Index.Type 		= tables["name"]
		indexed.Index.ID 		= strconv.Itoa(int(articleElastics.ID))

		indexing, err := json.Marshal(indexed)
		if err != nil {
			return false, err
		}

		articles, err := json.Marshal(articleElastics)
		if err != nil {
			return false, err
		}

		articleLists += string(indexing) + "\n"
		articleLists += string(articles) + "\n"

	}

	if err = rows.Err(); err != nil {
		return false, err
	}

	result, err := insertDataBulking(articleLists, tables["name"])
	if err != nil {return result, err}

	return true, nil
}

func insertDataBulking(data string, table string) (bool, error) {
	var responseBody map[string]interface{}
	var url = "http://" + config.GetString("database.elasticsearch.url") +"/"+ config.GetString("database.elasticsearch.prefix") +table+ "/_bulk?pretty=true"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return false, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &responseBody)
	if err != nil {return false, err}

	if reflect.ValueOf(responseBody["errors"]).Bool() {
		var errMsg interface{}
		index := reflect.ValueOf(responseBody["items"]).Index(0)

		for _, key := range reflect.ValueOf(index.Interface()).MapKeys() {
			strc := reflect.ValueOf(index.Interface()).MapIndex(key)

			if reflect.ValueOf(strc.Interface()).Kind() == reflect.Map {
				for _, key := range reflect.ValueOf(strc.Interface()).MapKeys() {
					strc2 := reflect.ValueOf(strc.Interface()).MapIndex(key)

					if key.Interface().(string) == "error" {
						errMsg = strc2.Interface()
					}
				}
			}
		}

		if errMsg == nil {
			errMsg = reflect.ValueOf(responseBody["item"]).Index(0).Interface()
		}

		err = fmt.Errorf("Something Error with Elasticsearch..\nError: %v", errMsg)
		return false, err
	}

	if strings.Split(response.Status, " ")[0] != strconv.Itoa(http.StatusOK) {
		status := fmt.Sprintf("response status insert articles : %v", response.Status)
		err = fmt.Errorf(status)
		return false, err
	}

	response.Body.Close()

	return true, nil
}

func getTagsArticle(db *sql.DB, id uint) ([]models.Tags, error) {
	var tags []models.Tags
	var sqlWord = " SELECT `tags`.`id`, `tags`.`title` FROM `articles_tags` "+
				  " LEFT JOIN `tags` ON `articles_tags`.`tag_id` = `tags`.`id`" +
				  " WHERE `articles_tags`.`article_id` = "

	rows, err := db.Query(sqlWord +strconv.Itoa(int(id)))
	if err != nil {
		return []models.Tags{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var tag models.Tags

		err = rows.Scan(
				&tag.ID,
				&tag.Title,
			)

		if err != nil {
			return []models.Tags{}, err
		}

		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return []models.Tags{}, err
	}

	return tags, nil
}

func getCategoriesArticle(db *sql.DB, id uint) ([]models.Categories, error) {
	var categories []models.Categories
	var sqlWord = " SELECT `categories`.`id`, `categories`.`title` FROM `articles_categories` "+
				  " LEFT JOIN `categories` ON `categories`.`id` = `articles_categories`.`category_id` "+
				  " WHERE `articles_categories`.`article_id` = "

	rows, err := db.Query(sqlWord +strconv.Itoa(int(id)))
	if err != nil {
		return []models.Categories{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Categories

		err := rows.Scan(
				&category.ID,
				&category.Title,
			)

		if err != nil {
			return []models.Categories{}, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return []models.Categories{}, err
	}

	return categories, nil
}

func getProductsArticle(db *sql.DB, id uint) ([]models.Products, error) {
	var products []models.Products
	var sqlWord = " SELECT `products`.`id`, `products`.`name`, `products`.`thumbnail`, `products`.`price`, `products`.`site_url`, `brands`.`id`, `brands`.`name` "+
				  " FROM `articles_products` LEFT JOIN `products` ON `products`.`id` = `articles_products`.`product_id` "+
				  " LEFT JOIN `brands` ON `brands`.`id` = `products`.`brand_id` "+
				  " WHERE `articles_products`.`article_id` = "

	rows, err := db.Query(sqlWord +strconv.Itoa(int(id)))
	if err != nil {
		return []models.Products{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Products

		err = rows.Scan(
				&product.ID,
				&product.Name,
				&product.Thumbnail,
				&product.Price,
				&product.SiteURL,
				&product.Brand.ID,
				&product.Brand.Name,
			)

		if err != nil {
			return []models.Products{}, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return []models.Products{}, err
	}

	return  products, nil
}