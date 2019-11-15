package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//creates a channel that will recieve a message when 30
	//s have eleapsed
	done := time.After(30 * time.Second)
	//make a new channel for passing byes for stdin to std
	//channel can only hold one mssage at a time as no size
	//has been specified
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		//use select to pass data from stdin to std out when recieved
		//or to shutdown when time out event occurs
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

//takes a write only channel chan <- and sends recieved input to that channel
func readStdin(out chan<- []byte) {
	for {
		//copies some data from stdin into data. note file.read blocks until it recieves data
		data := make([]byte, 1024)
		x, _ := os.Stdin.Read(data)
		if x > 0 {
			//sends buffered data over channel
			out <- data
		}
	}
}

// note on pausing with sleep

// func main() {
// 	//blocks for five seconds
// 	time.Sleep(5 * time.Second)

// 	//creates a channel that gers notiveid in five seconds and then
// 	//blocks until that channel recieves a notification
// 	sleep := time.After(5 * time.Second)
// 	<-sleep
// }
