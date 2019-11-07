package main

//paas cloud services - using config files as source codebase.
//config oppertunties are limited to things like environmental variables
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	fmt.Println(homePage)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//gets the port from the environment
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
