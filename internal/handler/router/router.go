package router

import (
	"context"
	"net/http"
	"regexp"
)

type RouterEntry struct {
	Method  string
	Path    *regexp.Regexp
	Handler http.HandlerFunc
}

type Router struct {
	routes []RouterEntry
}

func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range sr.routes {
		params := e.Match(r)
		if params == nil {
			continue
		}
		ctx := context.WithValue(r.Context(), "params", params)
		e.Handler.ServeHTTP(w, r.WithContext(ctx))
		return
	}
	// for _, e := range sr.routes {
	// 	match := e.Match(r)
	// 	if !match {
	// 		continue
	// 	}
	// 	e.Handler.ServeHTTP(w, r)
	// 	return
	// }

	http.NotFound(w, r)
}

func (r *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	exactPath := regexp.MustCompile("^" + path + "$")
	e := RouterEntry{
		Method:  method,
		Path:    exactPath,
		Handler: handlerFunc,
	}
	r.routes = append(r.routes, e)
}

// func readPathAndKeys(pattern string) *regexp.Regexp {
// 	paths := strings.Split(pattern, "/")
// 	for i, v := range paths {
// 		if strings.HasPrefix(v, ":") {

// 		}
// 	}
// 	return nil
// }

func (ro *RouterEntry) Match(r *http.Request) map[string]string {
	match := ro.Path.FindStringSubmatch(r.URL.Path)
	if match == nil {
		return nil
	}
	params := make(map[string]string)
	groupNames := ro.Path.SubexpNames()
	for i, group := range match {
		params[groupNames[i]] = group
	}
	// if r.Method != ro.Method {
	// 	return false
	// }
	// if r.URL.Path != ro.Path {
	// 	return false
	// }
	return params
}
