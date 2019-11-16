package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1) //make a buffered channel with one space
	for i := 1; i < 7; i++ {
		go worker(i, lock)
		//sets upt sex goroutines sharing locking channel
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	//a worker aquires lock by sending it a message. First worker to hit this will
	//get the one space and thus own the lock. the rest will blcok
	fmt.Printf("%d has teh lock\n", id)
	time.Sleep(500 * time.Millisecond)
	//the space between teh lock <-true and teh <-lock is 'locked'
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock
	//releases lock by reading a value which then opens that one space on teh buffer so the
	//next function can lock it
}
