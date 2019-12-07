package main

import (
	"io"
	"testing"
)

//interface verification -
//exporting types that implement external interfaces
//create interface types that describe external types
//rely on external interfaces and those change
//when use of that interface is restricted to type assertions

//write a 'canary't test that fails quickly if you make a mistake in your
//interface definition
//scenario: writing a customer wrieer that implements io.writer
//exporting this in library so other code uses it

//test the struct
type MyWriter struct {
	//....
}

func (m *MyWriter) Write([]byte) error {
	//write data somewhere
	return nil
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}
}

func doSomething(m map[string]interface{}) {
	w := m["w"].io.Writer //generates a runtime exception
}

//test to see if io is implemented
func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}

// if you run it iyou gett hs - the test files because Write method
// does not match sig of io.Writers Write([]byte) int error
// *MyWriter does not implement io.Writer (wrong type for Write method)
// have Write([]byte) error
// want Write([]byte) (int, error)
// FAIL	command-line-arguments [build failed]
