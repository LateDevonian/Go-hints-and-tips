package main

import (
	"bufio"
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
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	//tries to read a line of data from connection
	if err != nil {
		fmt.Println("failed to read from socket")
		conn.Close()
		//if fails to read, prints error and closes connection
	}
	response(data, conn)
	//given a line of text, passes it to a respsone
}

func response(data []byte, conn net.Conn) {
	//writes the data back out to the societ, echoing to client, close connection
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
}
