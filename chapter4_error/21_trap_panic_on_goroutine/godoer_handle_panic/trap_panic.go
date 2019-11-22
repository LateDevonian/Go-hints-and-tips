package main

import (
	"errors"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

//trap the panic and handle it so it does not crash the server

func message() {
	println("inside goroutine")
	panic(errors.New("oops"))
}

func main() {
	safely.Go(message)
	println("outside goroutines")
	time.Sleep(1000)
}
