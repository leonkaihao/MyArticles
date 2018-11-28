package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

//Open to init a db source
func (db *Database) Open(dbPath string) (err error) {
	log.Println("Open database ", dbPath, "...")
	dbObj, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic(err.Error)
	}
	db.DB = dbObj
	db.Path = dbPath
	err = db.createArticleTable()
	if err != nil {
		panic(err.Error)
	}
	err = db.createArticleTagTable()
	if err != nil {
		panic(err.Error)
	}
	// err = db.createTagTable()
	// if err != nil {
	// 	panic(err.Error)
	// }

	return
}

//Close a database source
func (db *Database) Close() (err error) {
	log.Println("Close database ", db.Path, "...")
	err = db.DB.Close()
	return
}

func (db *Database) createArticleTable() (err error) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS articles(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title VARCHAR(256) NULL,
        body VARCHAR(4096) NULL,
        date VARCHAR(16) NULL
    );
	`
	_, err = db.DB.Exec(sqlTable)
	return
}

func (db *Database) createArticleTagTable() (err error) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS articleTag(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        article_id INTEGER NULL,
        tag_name VARCHAR(64) NULL
    );
	`
	_, err = db.DB.Exec(sqlTable)
	return
}

//deprecated
func (db *Database) createTagTable() (err error) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS tags(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NULL
	);
	`
	_, err = db.DB.Exec(sqlTable)
	return
}
