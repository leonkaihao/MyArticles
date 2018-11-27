package database

import (
	"log"
)

//CreateArticle ...
func (db *Database) CreateArticle(article *ArticleRequest) (err error) {

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
