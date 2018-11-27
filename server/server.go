package server

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/leonkaihao/myarticles/handler/http/articles"
	"github.com/leonkaihao/myarticles/services/db"
)

//AppServer todo
type AppServer struct {
}

//Serve todo
func (appSvr *AppServer) Serve() error {
	err := db.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}
	artObj := new(articles.Articles)
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/articles", artObj.Create),
		rest.Get("/articles/:id", artObj.ArticleByID),
		rest.Get("/tags/:tagName/:date", artObj.ArticlesByTagDate),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	api.SetApp(router)
	addr := ":3000"
	log.Println("Start server", addr, "...")
	log.Fatal(http.ListenAndServe(addr, api.MakeHandler()))
	return err
}
