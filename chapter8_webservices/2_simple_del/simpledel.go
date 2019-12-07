package main

import (
	"fmt"
	"net/http"
)

func main() {
	//creates request object for delete http method
	req, _ := http.NewRequest("DELETE", "http//example.com/foo/bar", nil)
	//performs requiest with default client
	res, _ := http.DefaultClient.Do(req)
	//displays status code from request
	fmt.Printf("%s", res.Status)
}
