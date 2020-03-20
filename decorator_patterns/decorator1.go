package main

//https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/
//first example

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

// coolFunc takes in a function
// as a parameter
func coolFunc(a func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	// it then immediately calls that functino
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)

	//call coolFunc function, pass in myfunc
	coolFunc(myFunc)
	//this wraps the original function without altering the impleementation
}
