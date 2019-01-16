package controllers

import (
	"github.com/jacky-htg/shonet-frontend/libraries"
	"log"
	"net/http"
	_ "reflect"
)

func SearchIndexHandler(w http.ResponseWriter, r *http.Request) {

	authUser, err := GetUserClient(r)
	if libraries.CheckError(err) {
		return
	}

	log.Printf("\n\n Hasil : %v", authUser)
}
