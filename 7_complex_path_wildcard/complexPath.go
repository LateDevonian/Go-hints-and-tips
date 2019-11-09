package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}

//creates new initaialized path resolver
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

//add paths to internal lookup
type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	//adds paths to internal lookup
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//constructs our method and path to check
	check := req.Method + " " + req.URL.Path
	//iterate over registered paths
	for pattern, handlerFunc := range p.handlers {
		// checks whether current path matches a registered one
		if ok, err := path.Match(pattern, check); ok && err == nil {
			//excecutes handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	//if no path matches, page wasn't found
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Bob Dopolina"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Bob Dopolina"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
