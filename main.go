package main

import (
	"go-router/packages/controllers"
	"go-router/packages/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router{}
	r.NewRoute("GET", `/home/(?P<id>\d+)`, controllers.Home)
	r.NewRoute("GET", `.*`, controllers.NotFound)

	http.HandleFunc("/", r.Serve)
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("./src/"))))

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
