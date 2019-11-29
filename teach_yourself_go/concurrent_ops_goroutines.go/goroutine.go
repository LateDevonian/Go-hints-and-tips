package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowFunc()
	// the first attempt at this just had this line which
	//showed that the call to slofucnc was never seen
	//fmt.Println("I am now shown straightaway")
	//so we add a sleep
	fmt.Println("i am not shown until slwoFunc() completes")
	time.Sleep(time.Second * 3)

}
