package database

import (
	"github.com/jmoiron/sqlx"
)

//Database to identify a database source
type Database struct {
	Path string
	DB   *sqlx.DB
}

//ArticleRequest is input in DAO interface for creating
type ArticleRequest struct {
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

//ArticleResponse is returned to http req
type ArticleResponse struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

//TagDateResponse is returned to http req
type TagDateResponse struct {
	TagName     string   `json:"tag"`
	Count       int64    `json:"count"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

type articleTagResult struct {
	ArticleID string `db:"article_id"`
	TagName   string `db:"tag_name"`
}

//ArticleDAO ...
type ArticleDAO struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Date  string `db:"date"`
	Body  string `db:"body"`
}

//ArticleTagDAO ...
type ArticleTagDAO struct {
	ID        int64  `db:"id"`
	ArticleID int64  `db:"article_id"`
	TagName   string `db:"tag_name"`
}
