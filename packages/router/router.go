package router

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	routes []route
}

// NewRouter creates new route and appends it to Router,
// method specifies the method that is allowed,
// regexp must contain a named group,
// parsed value will be accessable through handler's context via GetField function
//
// regexp example -> https://regex101.com/r/84S9iL/1
func (router *Router) NewRoute(method, regexpString string, handler http.HandlerFunc) {
	regex := regexp.MustCompile("^" + regexpString + "$")
	method = strings.ToUpper(method)

	router.routes = append(router.routes, route{
		method,
		regex,
		handler,
	})
}

// If there is a exact match Serve redirects to the associated handler function.
//
// When matched, regular expression groups are used as key value pairs
// accessible in handler's context via GetField function
func (router *Router) Serve(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		match := route.regex.FindStringSubmatch(r.URL.Path)
		if len(match) > 0 {
			if strings.ToUpper(r.Method) != route.method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			matchMap := make(map[string]string)
			groupName := route.regex.SubexpNames()

			// map group name(key) to submatched result
			// these arrays have one to one relationship
			for i := 1; i < len(match); i++ {
				matchMap[groupName[i]] = match[i]
			}
			ctx := context.WithValue(r.Context(), struct{}{}, matchMap)
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
}

// Gets key value pairs from matched URL variables
func GetField(r *http.Request, name string) string {
	fields := r.Context().Value(struct{}{}).(map[string]string)
	return fields[name]
}
