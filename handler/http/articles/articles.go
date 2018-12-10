package articles

import (
	"net/http"
	"strconv"
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

	//ignore data Date field from client, and generate time from server
	data.Date = time.Now().Format("2006-01-02")
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
	idStr := r.PathParam("id")

	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := art.DB.GetArticleByID(id64)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(resp)
	return
}

//ArticlesByTagDate Get articles by a tag and date
//get artcles by tags and date range
func (art *Articles) ArticlesByTagDate(w rest.ResponseWriter, r *rest.Request) {
	tagName := r.PathParam("tagName")
	date := r.PathParam("date")

	resp, err := art.DB.GetArticlesByTagDate(tagName, date)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.WriteJson(resp)
	return
}

//Articles ...
func (art *Articles) Articles(w rest.ResponseWriter, r *rest.Request) {
	// keys := r.Form.Get("Keywords")
	// dateStart := r.Form.Get("StartFrom")
	// dateEnd := r.Form.Get("EndTo")

	return
}

//DeleteArticleByID ...
func (art *Articles) DeleteArticleByID(w rest.ResponseWriter, r *rest.Request) {

	return
}

//UpdateArticleByID ...
func (art *Articles) UpdateArticleByID(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	r.URL.Query()
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := &database.ArticleRequest{}
	err = r.DecodeJsonPayload(data)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = art.DB.UpdateArticle(idInt, data)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
