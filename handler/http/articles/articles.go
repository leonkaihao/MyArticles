package articles

import (
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

//Articles This class collect article operations
type Articles struct {
	sync.RWMutex
}

//Create ...
//create an article
func (art *Articles) Create(w rest.ResponseWriter, r *rest.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.WriteJson(map[string]string{"message": "not implement"})
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
