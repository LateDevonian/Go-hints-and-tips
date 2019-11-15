package main

import (
	"fmt"
	"runtime"
)

//this short example shows how a goroutine works, don't use it in the wild
//as it's likely the go function won't have time to run in a real world.
func main() {
	fmt.Println("Outside a goroutine")
	//runs first
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again")
	//runs second

	runtime.Gosched()
	//sets of the go function by giving it the oppertunity to run
	//the other goroutine before it exists
	//if you  get rid of this line it never runs the go as it's done
	//before the oppertunity to run it.
}
