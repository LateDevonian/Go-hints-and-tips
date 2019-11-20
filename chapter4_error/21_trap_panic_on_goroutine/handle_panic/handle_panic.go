package main

//same data as echo server with a non crash on  panic.

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

//run this by typing
//telnet localhost 1026
//sends connection and termination instructions
//then type test. will echo and close

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	//starts new server listening on port 1026
	if err != nil {
		fmt.Println("Failed to pen port on 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		//listens for new client connections and handles connection errors
		if err != nil {
			fmt.Println("error accepting connection")
			continue
		}
		go handle(conn)
		//when a connection accepted, passes it to handle function

	}
}
func handle(conn net.Conn) {
	//new deferred function handles the panic and makes sure
	//that in all cases the connection is closed
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("fatal error : %s", err)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket")
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("pretend i'm a real error"))
	//issue a panic to simulate a failure
}
