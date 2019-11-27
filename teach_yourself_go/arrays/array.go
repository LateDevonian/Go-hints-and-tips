package main

import (
	"fmt"
)

//array needs to be initalized with size
//different to ruby!
// var cheeses [2] string - this is a an array length of two strings
//can now assign a string to the elements
//cheeses[0] = "edam"
//bracket after variable is position.
//to print values of an array call it same as ruby fmt.Println(cheeses[0])
//to print it all just write cheeses.

func main() {
	//make an array
	var cheeses [2]string
	cheeses[0] = "edam"
	cheeses[1] = "blue"
	fmt.Println(cheeses)
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	//you can't add or change arrays so they kind of suck

}
