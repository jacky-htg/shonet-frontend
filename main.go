package main

import (
	"log"
	"net/http"

	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/routing"
)

func main() {
	router := routing.NewRouter()
	err := http.ListenAndServe(config.GetString("server.address"), router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
