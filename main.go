package main

import (
	"log"

	"github.com/leonkaihao/myarticles/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	appServer := new(server.AppServer)
	err := appServer.Serve()
	if err != nil {
		log.Fatalln("Failed to start APP Server:", err)
	}
}
