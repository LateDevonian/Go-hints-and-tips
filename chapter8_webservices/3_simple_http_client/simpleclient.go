package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	//create custom http climeout with timeout of 1 second
	cc := &http.Client{Timeout: time.Second}
	//handles timeout
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
