package main

import (
	"fmt"
	"time"
)

//to handle channels that don't close you want to use an extra
//channel to indicate that you're done. then the channel is closed
//cleanly

func main() {
	msg := make(chan string)
	//add a boolean channel to indicate when finished
	done := make(chan bool)
	until := time.After(5 * time.Second)

	//pass two channels into send
	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	//ch recieves, done sends
	for {
		select {
		case <-done:
			fmt.Println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
