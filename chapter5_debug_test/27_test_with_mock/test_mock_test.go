package main

import (
	"fmt"
	"testing"
)

//the method and struct one needs to mock in the code
//doesn't matter what it is
type Message struct {
	//
}

func (m *Message) Send(email, subject string, body []byte) error {
	//
	return nil
}

func main() {
	fmt.Println("this is the main message")

}

//use an interface to mock out the method

type Messager interface {
	//define an interface that describes teh methods you use on Message
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	//passes the interface instead of the message type
	return m.Send("noc@example.com", "critical error", problem)
}

//because we have created an abstraction from Message to Messenger
//can write a mock and use that in testing
type MockMessage struct {
	email, subject string
	body           []byte
}

//mockmessage implements messenger
func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	//create new mockmessage
	msgr := new(MockMessage)
	body := []byte("critical error")

	//runs alert method with mock
	Alert(msgr, body)

	if msgr.subject != "critical error" {
		//access mockmessage preoblerpies to verify results
		t.Errorf("Expected critical error, got '%s'", msgr.subject)
	}
}
