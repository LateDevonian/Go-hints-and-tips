package main

import (
	"fmt"
	"time"
)

//blocking code. line 15 won't print until call to slowfunc() done
func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("slseeper finished")
}

func main() {
	slowFunc()
	fmt.Println("I am not shown until slowFunc() completed")
}
