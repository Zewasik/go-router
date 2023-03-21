package main

import (
	"log"
	"net/http"

	httpRouting "github.com/Zewasik/go-router"
)

func main() {
	r := httpRouting.NewRouter()
	c := httpRouting.CORS{
		Origin:      "http://localhost:3000",
		Headers:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
		Methods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		Credentials: true,
	}

	r.NewRoute("get", `/home/(?P<id>\d+)`, Home)
	r.NewRoute("GET", `.*`, NotFound)

	http.HandleFunc("/", r.ServeWithCORS(c))

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	id, err := httpRouting.GetFieldString(r, "id")
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
