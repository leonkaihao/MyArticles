package database

import (
	"github.com/jmoiron/sqlx"
)

//Database to identify a database source
type Database struct {
	Path string
	DB   *sqlx.DB
}

//ArticleRequest will be input in DAO interface for creating
type ArticleRequest struct {
	Title string   `json: "title"`
	Date  string   `json: "date"`
	Body  string   `json: "body"`
	Tags  []string `json: "tags"`
}

type articleDAO struct {
	Title string `db:"title"`
	Date  string `db:"date"`
	Body  string `db:"body"`
}

type articleTagDAO struct {
	ArticleID int64  `db:"article_id"`
	TagName   string `db:"tag_name"`
}
