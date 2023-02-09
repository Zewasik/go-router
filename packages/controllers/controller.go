package controllers

import (
	"go-router/packages/httpRouting"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id := httpRouting.GetField(r, "id")

	w.Write([]byte("Home page, id: " + id))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Page not found"))
}
