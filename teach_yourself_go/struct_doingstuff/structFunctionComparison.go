package main

import (
	"fmt"
	"reflect"
)

type Alarm struct {
	Time  string
	Sound string
	Set   bool
}

func NewAlarm(time string) Alarm {
	//set default using a fuction that you call and pass the time
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
		Set:   true,
	}
	return a
}

//compare structs here
type Drink struct {
	Name string
	Ice  bool
}

func main() {
	//set default value using a fuction
	fmt.Printf("%+v\n", NewAlarm("07:00"))

	//compare
	dr := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	dri := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	if dr == dri {
		fmt.Println("both drinks are teh same")
	}
	fmt.Printf("%+v\n", dr)
	fmt.Printf("%+v\n", dri)

	fmt.Println("use reflect package to check types are the same")

	alarm := NewAlarm("07:00")
	fmt.Println(reflect.TypeOf(dr))
	fmt.Println(reflect.TypeOf(dri))
	fmt.Println(reflect.TypeOf(alarm))
}
