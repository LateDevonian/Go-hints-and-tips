package main

//an interface is a method set
//describes all the methods in a set but does not implement them
//this is a hodgepodge but interfaces are starting to make sense now
//the point is that an interface can take anyting and change. like
//what if youwant to swap robot for android. you don't need to
//swap code that refers to robot/android you just change what goes into the interfce
import (
	"errors"
	"fmt"
)

type Robot interface {
	//interface takes the powerOn method
	PowerOn() error
}

type T850 struct {
	Name string
}

//poweron method - takes no arguements, returns an error
//takes a struct
func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}
func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}
	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
