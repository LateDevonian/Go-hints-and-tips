//get a stack trace using the runtime tool

package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)
}

//simple stack debug is below

// import (
// 	"runtime/debug"
// )

// func main() {
// 	foo()
// }

// func foo() {
// 	bar()
// }

// func bar() {
// 	debug.PrintStack()
// }
