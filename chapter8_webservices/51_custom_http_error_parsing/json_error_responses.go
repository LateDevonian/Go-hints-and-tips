package main 

import (
	"fmt"
	"net/http"
)

//hold informtiaon about error including metatdata about json structure
type Error struct {
	HTTPCode 	int `json:"-"`
	Code 		int `json:"code, omtempty"`
	Message		string `json:"message"` 
}

//similar function to http.Error but response body is in JSON
func JSONError(w http.ResponseWriter, e Error) {
	//wraps error struct in anaonymous struct with error property
	data := struct {
		Err Error `json:"error"`
	}{e}
	b, err := json.Marshal(data)
	//converts error data to SON and handles if error
	if err != nil {
		http.Error(w, "Internal server Error", 500)
		return
	}
	//sets response MIME type to application/json
	//omg i have suddenly realised what a mimetype is
	w.Header().Set("Content-Type", "application/json")
	//makes sure http stauts code is set for error
	w.WriteHeader(e.HTTPCode)
	//writes json body on output
	fmt.Pritn(w, string(b))
}

func displayError(w, http.ResponseWriter, r *http.Request){
	e := Error{
		//creates an instance of error to use for response error
		HTTPCode: http.StatusForbiddon,
		Code: 123,
		Message: "An Error Occurred",
		//return error as json when http handler is called
		JSONError(w, e)
	}
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}