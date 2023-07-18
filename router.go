package httpRouting

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type router struct {
	routes      map[string][]route
	corsHeaders map[string]string
}

type ContextKey string

func NewRouterBuilder() *router {
	return &router{
		routes:      make(map[string][]route),
		corsHeaders: make(map[string]string),
	}
}

func (r *router) SetAllowOrigin(o string) *router {
	o = strings.TrimSpace(o)
	if o != "" {
		r.corsHeaders["Access-Control-Allow-Origin"] = o
	}
	return r
}

func (r *router) SetAllowMethods(m []string) *router {
	if len(m) > 0 {
		r.corsHeaders["Access-Control-Allow-Methods"] = strings.Join(m, ", ")
	}
	return r
}

func (r *router) SetAllowHeaders(h []string) *router {
	if len(h) > 0 {
		r.corsHeaders["Access-Control-Allow-Headers"] = strings.Join(h, ", ")
	}
	return r
}

func (r *router) SetCredantials(c bool) *router {
	if c {
		r.corsHeaders["Access-Control-Allow-Credentials"] = "true"
	}
	return r
}

// NewRouter creates new route and appends it to Router,
// method specifies the method that is allowed,
// regexp must contain a named group,
// parsed value will be accessable through handler's context via GetField function
//
// regexp example -> https://regex101.com/r/84S9iL/1
func (r *router) NewRoute(method, regexpString string, handler http.HandlerFunc, middlewareBefore ...Middleware) {
	regex := regexp.MustCompile("^" + regexpString + "$")
	method = strings.ToUpper(method)

	r.routes[method] = append(r.routes[method], route{
		regex,
		handler,
		middlewareBefore,
		make([]Middleware, 0),
	})
}

// If there is a exact match Serve redirects to the associated handler function.
//
// When matched, regular expression groups are used as key value pairs
// accessible in handler's context via GetField function
func (r *router) Serve(w http.ResponseWriter, req *http.Request) {
	r.serve(w, req)
}

// Same usage as in Serve function, but also adds specified CORS headers
func (r *router) ServeWithCORS() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for header, value := range r.corsHeaders {
			w.Header().Set(header, value)
		}

		if req.Method == "OPTIONS" {
			return
		}

		r.serve(w, req)
	}
}

func (r *router) serve(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes[strings.ToUpper(req.Method)] {
		match := route.regex.FindStringSubmatch(req.URL.Path)
		if len(match) > 0 {
			matchMap := make(map[string]string)
			groupName := route.regex.SubexpNames()

			// map group name(key) to submatched result
			// these arrays have one to one relationship
			for i := 1; i < len(match); i++ {
				matchMap[groupName[i]] = match[i]
			}
			ctx := context.WithValue(req.Context(), ContextKey("requestParams"), matchMap)
			req = req.WithContext(ctx)

			handler := route.handler
			for i := len(route.middlewareBefore) - 1; i >= 0; i-- {
				handler = route.middlewareBefore[i](handler)
			}

			handler.ServeHTTP(w, req)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

// Returns the string value of the given key from matched URL variables
func GetRequestParamString(r *http.Request, name string) (string, error) {
	fields, ok := r.Context().Value(ContextKey("requestParams")).(map[string]string)
	if !ok {
		return "", fmt.Errorf("internal error: no fileds in context")
	}

	field, exist := fields[name]
	if !exist {
		return "", fmt.Errorf("no such variable in request: %v", name)
	}

	return field, nil
}

// Returns the integer value of the given key from matched URL variables
func GetRequestParamInt(r *http.Request, name string) (int, error) {
	field, err := GetRequestParamString(r, name)
	if err != nil {
		return 0, err
	}

	fieldInt, err := strconv.Atoi(field)
	if err != nil {
		return 0, err
	}

	return fieldInt, nil
}
