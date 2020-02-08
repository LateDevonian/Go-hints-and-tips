package main 

import { 
	"encoding/json"
	"fmt"
}
//struct represents information in JSON. json tab maps name property to name in json
type Person struct {
	Name string `json:"name"`
}
//represent json as a string
var JSON = `{
	"name": "Bob Dopolina"
	}`
 func main() {
	 //hold the parsed json data
	 var p Person
	 //parse json ito the person struct
	 err := json.Unmarshal([]byte(JSON), &p)
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 //acts on person object and prints it
	 fmt.Println(p)
 }