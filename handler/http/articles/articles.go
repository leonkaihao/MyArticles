package articles

import (
	"net/http"
	"sync"

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
	w.WriteHeader(http.StatusForbidden)
	w.WriteJson(map[string]string{"message": "not implement"})
	return
}

//ArticlesByTagDate Get articles by a tag and date
//get artcles by tags and date range
func (art *Articles) ArticlesByTagDate(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.WriteJson(map[string]string{"message": "not implement"})
}
