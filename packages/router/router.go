package router

import (
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes []route
}

// creating new route and appending it to router
//
// regexp example -> https://regex101.com/r/84S9iL/1
func (router *Router) NewRoute(method, regexpString string, handler http.HandlerFunc) {
	regex := regexp.MustCompile("^" + regexpString + "$")
	router.routes = append(router.routes, route{
		method,
		regex,
		handler,
	})
}
