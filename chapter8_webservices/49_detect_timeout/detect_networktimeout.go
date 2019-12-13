package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//check to see if the error returned is a timeout in your client
//net will have some timeout errors but
//it's good to check. so add has timed out function

func main() {
	//create custom http climeout with timeout of 1 second
	cc := &http.Client{Timeout: time.Second}
	//handles timeout
	res, err := cc.Get("http://example.com/test.zip")
	if err != nil && hasTimedOut(err) {
		fmt.Println("a timeout error occured")
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
