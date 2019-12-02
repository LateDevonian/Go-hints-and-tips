//buffered channels

package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	//initialise a buffered channel with a length of 2
	messages := make(chan string, 2)
	//send two messages to the channel
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("pushed two messages onto channel with no recievers")
	time.Sleep(time.Second * 1)
	//the channel is passed as an argument to a function called reciever, which iterates over the channel
	//using range and prints buffereed messages in channel to console
	receiver(messages)
}
