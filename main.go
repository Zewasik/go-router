package main

import (
	"go-router/controllers"
	"go-router/packages/httpRouting"
	"log"
	"net/http"
)

func main() {
	r := httpRouting.NewRouter()

	r.NewRoute("get", `/home/(?P<id>\d+)`, controllers.Home)
	r.NewRoute("GET", `.*`, controllers.NotFound)

	http.HandleFunc("/", r.ServeWithCORS)

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
