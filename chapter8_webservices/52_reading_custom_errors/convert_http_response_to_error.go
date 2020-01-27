package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code, omitempty"`
	Message  string `json:"message"`
}

func (e Error) Error() string {
	fs := "HTTP: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

//use this Get instead of http.Get - uses http.Get but handles it differently
func get(u string) (*http.Response, error) {
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		//chekck otu reponse content type and error if wrong
		//check outsie 200 code
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "Unknown error. HTTP Status: %s"
		}
		b, _ := ioutil.ReadAll(res.Body)
		//reads body of response into a buffer
		res.Body.Close()
		var data struct {
			//parse JSON respons and puts data into a struct
			//repsonds to errors
			Err Error `json:"error"`
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "Unable to parse json: %s. HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		data.Err.HTTPCde = res.StatusCode
		//adds http status code to error and returns custom error
		//and response
		return res, data.Err
	}
	return res, nil
}

//the Error() method applied to the Error type inmlements the error
//interface. allows instances of Error to be passed between fucntions
//as an error

func main() {
	res, err := ger("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
