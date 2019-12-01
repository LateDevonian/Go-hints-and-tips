package main

import (
	"log"
	"net"
	"time"
)

//run nc -lk 1902 to satart a server using netcat
//tcp logging is prone to back pressure but UDP
//logging can't guarantee data accuracy

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	logger.Println("this is a regular message.")
	logger.Panicln("This is a panic")
}
