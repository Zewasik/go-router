package httpRouting

import (
	"net/http"
	"regexp"
)

type route struct {
	regex            *regexp.Regexp
	handler          http.Handler
	middlewareBefore []Middleware
	middlewareAfter  []Middleware
}
