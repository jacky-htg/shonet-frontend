package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
)

type HotArticlesMenu struct {
	Articles 	 []models.Article
	HotArticles  []models.HotArticle
}

type TopProductMenu struct {
	Top			[]models.Product
	Products    []models.Product
}

var db *sql.DB
var err error

func init() {
	//Create an sql.DB and check for errors
	db, err = sql.Open(config.GetString("database.mysql.driverName"), config.GetString("database.mysql.dataSourceName"))
	libraries.CheckError(err)

	//test the connection
	err = db.Ping()
	libraries.CheckError(err)
}
