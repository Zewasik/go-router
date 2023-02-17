package main

import (
	"go-router/controllers"
	"go-router/packages/httpRouting"
	"log"
	"net/http"
)

func main() {
	r := httpRouting.NewRouter()
	c := httpRouting.CORS{
		Origin:      "http://localhost:3000",
		Headers:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		Methods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		Credentials: true,
	}

	r.NewRoute("get", `/home/(?P<id>\d+)`, controllers.Home)
	r.NewRoute("GET", `.*`, controllers.NotFound)

	http.HandleFunc("/", r.ServeWithCORS(c))

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
