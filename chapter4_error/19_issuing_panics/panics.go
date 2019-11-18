package main

// simple defer - recovery from a panic

//uses a deferred function: go gaurantees excectuion of a function
//the moment it's parent function returns

import "fmt"

func main() {
	defer goodbye()
	fmt.Println("hello world")
}

func goodbye() {
	fmt.Println("goodbye")
}

// technique 1 - runs and panics a bit useless
//func main() {
// 	panic(nil)
// }

//technique 2 - proper panic, idiomatic to pass an error to a panic
//func main() {
//	panic(errors.New("something bad"))
//}
