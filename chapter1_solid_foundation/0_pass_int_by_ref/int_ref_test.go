package main

import (
	"testing"
)

///failing

// type Student struct {
// 	ID   string
// 	Name string
// }

// type Teacher struct {
// 	ID   string
// 	Name string
// }

var stud Student
var tech Teacher

func TestDoSomethingWithParam(t *testing.T) {
	t.Log("send a student through")
	returned := doSomethingWithParam(&stud)

	if returned.ID != "124" {
		//access mockmessage preoblerpies to verify results
		t.Errorf("Expected critical error, got '%s'", returned.ID)
	}
}
