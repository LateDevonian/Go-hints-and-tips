package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		//provide a deferred closure to handle panic recovery
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	//call a fucntion that panics
	yikes()
}

//to catch the panic yikes raises, rewite a deferred closure function that
//checks for a panic and recovers if it finds one

func yikes() {
	//emits a panic with an error for a body
	panic(errors.New("something bad happened"))
}

// Scope for deferred closures:
// the closure may reference msg as it's defined before the cloure
// func main() {
// 	var msg string //define variable outside the closure
// 	defer funct() {
// 		fmt.Println(msg) //print variable inside deferred closure
// 	}()
// 	msg = "hello world" //set value of variable
// }

//msg out of scope, this returns a compile error
//message isnt' declared prior to def function, so msg is undefined
// func main() {
// 	defer func() {
// 		fmt.Println(msg)
// 	}()
// 	msg := "hello world"
// }
