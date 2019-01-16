package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
	"strings"
)

func GetArticlesById(ids []uint) ([]models.Article, error) {
	var idString []string
	for _, id := range ids { idString = append(idString, string(id)) }

	sql := "SELECT * FROM `articles` WHERE `articles`.`id` IN (" +strings.Join(idString, ", ")+ " ORDER BY `articles`.`publish_date` DESC"
	articles, err := fetchArticles(db.Query(sql))
	if err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func GetOneArticleById(id uint) (models.Article, error) {
	sql := "SELECT * FROM `articles` WHERE `articles`.`id` = " + string(id) + " AND `articles`.`status` = 'P'"
	articles, err := fetchArticles(db.Query(sql))
	if err != nil {
		return models.Article{}, err
	}

	return articles[0], nil
}

func GetHotArticles() (HotArticlesMenu, error) {

	var hotArticleSql string
	var result HotArticlesMenu
	var whereNotIn []string

	limitOrder := 4
	curLimit   := 0

	hotArticleSql = " SELECT * FROM `hot_articles` WHERE (`start_date` <= NOW() AND `end_date` >= NOW()) AND EXISTS(" +
				 	" SELECT * FROM `articles` WHERE `hot_articles`.`article_id` = `articles`.`id` AND (" +
				 	" `articles`.`status` = 'P' AND `articles`.`publish_date` <= NOW() " +
				 	" )" +
				 	" ) ORDER BY `hot_articles`.`id` DESC LIMIT " + string(limitOrder)


	hotArticles, err := fetchHotArticles(db.Query(hotArticleSql))
	if err != nil {
		return HotArticlesMenu{}, err
	}

	if len(hotArticles) > 0 {
		result.HotArticles = hotArticles

		for _, val := range hotArticles {
			whereNotIn = append(whereNotIn, string(val.ID))
		}

		curLimit += len(hotArticles)
	}

	articleSql   := " SELECT * FROM `articles` WHERE `articles`.`id` NOT IN (" +
					  strings.Join(whereNotIn, ", ")+
					" ) LIMIT " +
					  (string(limitOrder - curLimit))+
					" ORDER BY `articles`.`trending_count`, `articles`.`publish_date` DESC"

	if curLimit < limitOrder {
		articles, err := fetchArticles(db.Query(articleSql))
		if err != nil {
			return HotArticlesMenu{}, err
		}

		result.Articles = articles
	}

	return result, nil
}

func fetchHotArticles(row *sql.Rows, err error) ([]models.HotArticle, error) {
	var hotArticles []models.HotArticle

	libraries.CheckError(err)
	if err != nil {
		return []models.HotArticle{}, err
	}

	defer row.Close()

	for row.Next() {
		var hotArticle models.HotArticle
		var hotArticleNull models.HotArticleNull

		err = row.Scan(
				&hotArticle.ID,
				&hotArticle.Article.ID,
				&hotArticle.StartDate,
				&hotArticle.EndDate,
				&hotArticleNull.CreatedAt,
				&hotArticleNull.UpdatedAt,
			)

		if err != nil {
			return []models.HotArticle{}, err
		}

		hotArticle.CreatedAt = hotArticleNull.CreatedAt.Time
		hotArticle.UpdatedAt = hotArticleNull.UpdatedAt.Time

		hotArticle.Article, err = GetOneArticleById(hotArticle.Article.ID)
		if err != nil {
			return []models.HotArticle{}, err
		}

		hotArticles = append(hotArticles, hotArticle)
	}

	if err = row.Err(); err != nil {
		return []models.HotArticle{}, err
	}

	return hotArticles, nil
}

func fetchArticles(rows *sql.Rows, err error) ([]models.Article, error) {
	var articles []models.Article

	libraries.CheckError(err)
	if err != nil {
		return []models.Article{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var article models.Article
		var articleNull models.ArticleNull

		err := rows.Scan(
				&article.ID,
				&article.Title,
				&article.Slug,
				&article.Permalink,
				&articleNull.Content,
				&articleNull.Image,
				&articleNull.ImageSource,
				&articleNull.SeoKeyword,
				&article.Type,
				&article.Status,
				&articleNull.RequestPublishDate,
				&articleNull.PublishDate,
				&article.Writer.ID,
				&article.Editor.ID,
				&articleNull.CreatedAt,
				&articleNull.UpdatedAt,
				&articleNull.TrendingCount,
				&articleNull.ContentManipulation,
			)

		if err != nil {
			return []models.Article{}, err
		}

		article.Content = articleNull.Content.String
		article.Image = articleNull.Image.String
		article.ImageSource = articleNull.ImageSource.String
		article.SeoKeyword = articleNull.SeoKeyword.String
		article.RequestPublishDate = articleNull.RequestPublishDate.Time
		article.PublishDate = articleNull.PublishDate.Time
		article.CreatedAt = articleNull.CreatedAt.Time
		article.UpdateAt = articleNull.UpdatedAt.Time
		article.TrendingCount = uint(articleNull.TrendingCount.Int64)
		article.ContentManipulation = articleNull.ContentManipulation.String

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return []models.Article{}, nil
	}

	return articles, nil
}
