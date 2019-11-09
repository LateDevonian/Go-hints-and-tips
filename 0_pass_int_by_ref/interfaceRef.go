package main

//https://hackernoon.com/today-i-learned-pass-by-reference-on-interface-parameter-in-golang-35ee8d8a848e

//pass by refrenece on interface parameter

import (
	"fmt"
)

func main() {
	var item Student
	var teach Teacher

	doSomethingWithParam(&item)
	fmt.Printf("%+v", item)
	doSomethingWithParam(&teach)
	fmt.Printf("%+v", teach)
}

type Student struct {
	ID   string
	Name string
}

type Teacher struct {
	ID   string
	Name string
}

//generic function accepts interface{} and does something inside it
//swich condition can do many things with it
func doSomethingWithParam(item interface{}) {
	switch v := item.(type) {
	case *Student:
		*v = Student{
			ID:   "124",
			Name: "Bob Dopolina",
		}
	case *Teacher:
		*v = Teacher{
			ID:   "Teacher1",
			Name: "Dr Jones",
		}
	}
}
