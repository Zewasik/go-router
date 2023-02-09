package controllers

import (
	"go-router/packages/httpRouting"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id, err := httpRouting.GetField(r, "id")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Home page, id: " + id))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Page not found"))
}
