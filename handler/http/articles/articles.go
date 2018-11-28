package articles

import (
	"net/http"
	"sync"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/leonkaihao/myarticles/services/database"
)

//Articles This class collect article operations
type Articles struct {
	sync.RWMutex
	DB *database.Database
}

//Create ...
//create an article
func (art *Articles) Create(w rest.ResponseWriter, r *rest.Request) {
	data := &database.ArticleRequest{}
	if err := r.DecodeJsonPayload(&data); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(data.Title) == 0 {
		rest.Error(w, "Title must exist.", http.StatusBadRequest)
		return
	}
	if len(data.Title) >= 256 {
		rest.Error(w, "Title length must be less than 255.", http.StatusBadRequest)
		return
	}
	if len(data.Body) == 0 {
		rest.Error(w, "Body must exist.", http.StatusBadRequest)
		return
	}
	if len(data.Body) >= 4096 {
		rest.Error(w, "Body length must be less than 4K.", http.StatusBadRequest)
		return
	}
	for _, tag := range data.Tags {
		if len(tag) >= 64 {
			rest.Error(w, "tag string length must be less than 63.", http.StatusBadRequest)
			return
		}
	}
	//ignore data Date field from client, and generate time from server
	data.Date = time.Now().Format("2016-09-22")
	//protect a transaction
	art.Lock()
	defer art.Unlock()
	err := art.DB.CreateArticle(data)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

//ArticleByID Get one article by its id
//get an article by id
func (art *Articles) ArticleByID(w rest.ResponseWriter, r *rest.Request) {
	rest.Error(w, "not implement", http.StatusBadRequest)
	return
}

//ArticlesByTagDate Get articles by a tag and date
//get artcles by tags and date range
func (art *Articles) ArticlesByTagDate(w rest.ResponseWriter, r *rest.Request) {
	rest.Error(w, "not implement", http.StatusBadRequest)
	return
}
