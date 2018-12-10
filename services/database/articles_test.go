package database

import (
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestDatabase_CreateArticle(t *testing.T) {
	db := &Database{}
	db.Open("./test.db")
	defer db.Close()

	type args struct {
		article *ArticleRequest
	}
	tests := []struct {
		name    string
		db      *Database
		args    args
		wantErr bool
	}{
		{
			name: "empty title",
			db:   db,
			args: args{
				article: &ArticleRequest{
					Title: "",
					Date:  "Date",
					Body:  "Body",
					Tags:  []string{"tag1"},
				},
			},
			wantErr: true,
		},

		{
			name: "empty body",
			db:   db,
			args: args{
				article: &ArticleRequest{
					Title: "Title",
					Date:  "Date",
					Body:  "",
					Tags:  []string{"tag1"},
				},
			},
			wantErr: true,
		},

		{
			name: "empty tag",
			db:   db,
			args: args{
				article: &ArticleRequest{
					Title: "",
					Date:  "Date",
					Body:  "Body",
					Tags:  []string{""},
				},
			},
			wantErr: true,
		},
		{
			name: "normal format",
			db:   db,
			args: args{
				article: &ArticleRequest{
					Title: "title",
					Date:  "Date",
					Body:  "Body",
					Tags:  []string{"science"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.CreateArticle(tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("Database.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_GetArticleByID(t *testing.T) {
	os.Remove("./test.db")
	db := &Database{}
	db.Open("./test.db")
	defer db.Close()
	req := &ArticleRequest{Title: "title", Body: "body", Date: time.Now().Format("2006-01-02"), Tags: []string{"science"}}
	db.CreateArticle(req)
	type args struct {
		id int64
	}
	tests := []struct {
		name     string
		db       *Database
		args     args
		wantResp *ArticleResponse
		wantErr  bool
	}{
		{
			name: "get first item",
			db:   db,
			args: args{
				id: 1,
			},
			wantResp: &ArticleResponse{
				ID:    1,
				Title: "title",
				Body:  "body",
				Date:  time.Now().Format("2006-01-02"),
				Tags: []string{
					"science",
				},
			},
			wantErr: false,
		},
		{
			name: "get second item",
			db:   db,
			args: args{
				id: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.db.GetArticleByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetArticleByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Database.GetArticleByID() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDatabase_GetArticlesByTagDate(t *testing.T) {

	os.Remove("./test.db")
	db := &Database{}
	db.Open("./test.db")
	defer db.Close()
	req := &ArticleRequest{Title: "title1", Body: "body", Date: time.Now().Format("2006-01-02"), Tags: []string{"science", "medical"}}
	db.CreateArticle(req)
	req = &ArticleRequest{Title: "title2", Body: "body", Date: time.Now().Format("2006-01-02"), Tags: []string{"science", "fitness"}}
	db.CreateArticle(req)
	req = &ArticleRequest{Title: "title3", Body: "body", Date: time.Now().Format("2006-01-02"), Tags: []string{"fitness", "sport"}}
	db.CreateArticle(req)

	type args struct {
		tagName string
		date    string
	}
	tests := []struct {
		name     string
		db       *Database
		args     args
		wantResp *TagDateResponse
		wantErr  bool
	}{
		{
			name: "tag date select 2 items",
			db:   db,
			args: args{
				tagName: "science",
				date:    time.Now().Format("2006-01-02"),
			},
			wantResp: &TagDateResponse{
				TagName: "science",
				Count:   2,
				Articles: []string{
					"1", "2",
				},
				RelatedTags: []string{
					"science", "medical", "fitness",
				},
			},
			wantErr: false,
		},
		{
			name: "tag date select 1 items",
			db:   db,
			args: args{
				tagName: "sport",
				date:    time.Now().Format("2006-01-02"),
			},
			wantResp: &TagDateResponse{
				TagName: "sport",
				Count:   1,
				Articles: []string{
					"3",
				},
				RelatedTags: []string{
					"fitness", "sport",
				},
			},
			wantErr: false,
		},
		{
			name: "tag not exist",
			db:   db,
			args: args{
				tagName: "child",
				date:    time.Now().Format("2006-01-02"),
			},
			wantErr: true,
		},
		{
			name: "date not exist",
			db:   db,
			args: args{
				tagName: "science",
				date:    "2006-01-02",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.db.GetArticlesByTagDate(tt.args.tagName, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.GetArticlesByTagDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Database.GetArticlesByTagDate() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestDatabase_UpdateArticle(t *testing.T) {

	os.Remove("./test.db")
	db := &Database{}
	db.Open("./test.db")
	defer db.Close()
	req := &ArticleRequest{Title: "title1", Body: "body", Date: time.Now().Format("2006-01-02"), Tags: []string{"science", "medical"}}
	db.CreateArticle(req)

	type args struct {
		id      int64
		article *ArticleRequest
	}
	tests := []struct {
		name    string
		db      *Database
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.UpdateArticle(tt.args.id, tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("Database.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
