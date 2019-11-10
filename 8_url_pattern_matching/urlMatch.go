package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//doesnt run, has nil pointer or memo. come back later

func main() {
	rr := newPathResolver()
	//registers paths to functions
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	http.ListenAndServe(":8080", rr)
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

//stores compiled regular expressuions for reuse
type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	//looks up and excecutes handler function
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
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
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Bob Dopolina"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
