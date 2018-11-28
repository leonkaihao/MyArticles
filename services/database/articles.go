package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
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
	itemArticle := &ArticleDAO{
		Title: article.Title,
		Date:  article.Date,
		Body:  article.Body,
	}
	schemaArticle := "insert into articles(title, date, body) values(:title, :date, :body)"
	res, err := db.DB.NamedExec(schemaArticle, itemArticle)
	if err != nil {
		return
	}
	//get the inserted item id
	articleID, err := res.LastInsertId()
	if err != nil {
		return
	}
	//and then insert all tags to articleTag table
	schemaArticleTag := "insert into articleTag(article_id, tag_name) values(:article_id, :tag_name)"
	for _, tag := range article.Tags {
		itemArticleTag := &ArticleTagDAO{
			ArticleID: articleID,
			TagName:   tag,
		}
		_, err = db.DB.NamedExec(schemaArticleTag, itemArticleTag)

		if err != nil {
			return
		}
	}

	return
}

//GetArticleByID get a article by its id
func (db *Database) GetArticleByID(id int64) (resp *ArticleResponse, err error) {

	schemaArticle := "SELECT * FROM articles WHERE id=?"
	article := &ArticleDAO{}
	err = db.DB.Get(article, schemaArticle, id)
	if err != nil {
		return
	}
	resp = &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
		Date:  article.Date,
	}
	schemaTags := "SELECT DISTINCT tag_name FROM articleTag WHERE article_id=?"
	err = db.DB.Select(&resp.Tags, schemaTags, id)
	if err != nil {
		return
	}

	return
}

//GetArticlesByTagDate get collect info by tag and date
func (db *Database) GetArticlesByTagDate(tagName string, date string) (resp *TagDateResponse, err error) {
	//get article IDs
	schemaArticleID := "SELECT articles.id FROM articles, articleTag WHERE articles.id=articleTag.article_id AND articleTag.tag_name=? AND articles.date=?"
	resp = &TagDateResponse{
		TagName: tagName,
	}
	err = db.DB.Select(&resp.Articles, schemaArticleID, tagName, date)
	if err != nil {
		return
	}

	resp.Count = int64(len(resp.Articles))
	if resp.Count == 0 {
		err = errors.New("No article match this tag/date")
		resp = nil
		return
	}
	//get tags
	schemaTagsID := "SELECT DISTINCT tag_name FROM articleTag WHERE article_id IN (?)"
	query, args, err := sqlx.In(schemaTagsID, resp.Articles)
	if err != nil {
		return
	}
	query = db.DB.Rebind(query)

	err = db.DB.Select(&resp.RelatedTags, query, args...)
	if err != nil {
		return
	}
	return
}
