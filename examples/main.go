package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httpRouting "github.com/Zewasik/go-router"
)

func main() {
	r := httpRouting.NewRouterBuilder().
		SetAllowOrigin("http://localhost:3000").
		SetAllowMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}).
		SetAllowHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token"}).
		SetCredantials(true)

	r.NewRoute("get", `/home/(?P<id>\d+)`, Home, HasAccess())

	http.HandleFunc("/", r.ServeWithCORS())

	log.Println("Ctrl + Click on the link: http://localhost:8080")
	log.Println("To stop the server press `Ctrl + C`")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Home(w http.ResponseWriter, r *http.Request) {
	id, err := httpRouting.GetRequestParamString(r, "id")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Home page, id: " + id))
}

func HasAccess() httpRouting.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := httpRouting.GetRequestParamInt(r, "id")
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if id != 1 {
				json.NewEncoder(w).Encode("you are not allowed to see this page")
				return
			}

			ctx := context.WithValue(r.Context(), httpRouting.ContextKey("user"), "aboba user 1")
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
