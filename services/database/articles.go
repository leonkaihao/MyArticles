package database

import (
	"errors"
	"log"
)

//CreateArticle ...
func (db *Database) CreateArticle(article *ArticleRequest) (err error) {
	if len(article.Title) == 0 {
		err = errors.New("Title must exist")
		return
	}
	if len(article.Title) >= 256 {
		err = errors.New("Title length must be less than 255")
		return
	}
	if len(article.Body) == 0 {
		err = errors.New("Body must exist")
		return
	}
	if len(article.Body) >= 4096 {

		err = errors.New("Body length must be less than 4K")
		return
	}
	for _, tag := range article.Tags {
		if len(tag) >= 64 {
			err = errors.New("Tag string length must be less than 63")
			return
		}
	}
	if len(article.Date) >= 16 {
		err = errors.New("Date string length must be less than 15")
		return
	}
	//first insert to articles table
	itemArticle := &articleDAO{
		Title: article.Title,
		Date:  article.Date,
		Body:  article.Body,
	}
	schemaArticle := "insert into articles(title, date, body) values(:title, :date, :body)"
	res, err := db.DB.NamedExec(schemaArticle, itemArticle)
	if err != nil {
		log.Fatalln("db.CreateArticle:", err)
	}
	//get the inserted item id
	articleID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("db.CreateArticle:", err)
	}
	//and then insert all tags to articleTag table
	schemaArticleTag := "insert into articleTag(article_id, tag_name) values(:article_id, :tag_name)"
	for _, tag := range article.Tags {
		itemArticleTag := &articleTagDAO{
			ArticleID: articleID,
			TagName:   tag,
		}
		_, err := db.DB.NamedExec(schemaArticleTag, itemArticleTag)

		if err != nil {
			log.Fatalln("db.CreateArticle:", err)
		}
	}

	return
}
