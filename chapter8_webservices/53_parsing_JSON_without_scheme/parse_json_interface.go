package main

//parse a json data into an interface instead of a struct if i don't know
//what structure it will be. then use it.
import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{ 
	"firstname": "Jean",
	"secondname": "Avril",
	"age": 46,
	"education": [
		{
			"institution": "monash universtiy",
			"degree": "BSC hons"
		},
		{ 
			"institution": "deakin university",
			"degree": "Grad cert in writing"
		}
	]
	}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f) //shoves ks into the interface
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f) //access the JSON data in the interface
}

//note when json is parsed into a struct, it's accessible. name will be p.Name
//get an error in an interface
//need to accepss it ias a type other than an interface{}
//in this case JSON is an object o you can use type map[]string interface, to provide
//access to the next level of data. so to access firstname:
//m := f.(map[string]interface{})
//fmt.Println(m["firstname"])
