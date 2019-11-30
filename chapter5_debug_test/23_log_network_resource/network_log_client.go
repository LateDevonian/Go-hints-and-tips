//"treat logs as event streams"
//
package main

import (
	"log"
	"net"
)

//you have to start a server
//use netcat
// nc -lk 1902
//run the go program and then check out what is
//happening in the network tab! it's logged!

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
