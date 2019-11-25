package main

import (
	"fmt"
)

//methods that take a pointer can modify data elements from
//an original struct, as they work from the same memory.
type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	//takes a pointer reference

	return 0.5 * (t.base * t.height)
}

func (t Triangle) changeBase(f float64) {
	//takes a value
	//no update when calling this method on intialised struct
	t.base = f
	return
}

func (t *Triangle) changeRefBase(f float64) {
	//takes reference so this will update an initalised struct
	t.base = f
	return
}

func main() {
	//initalised struct, does the calc
	// asterisk before Triangle: reciever is pointer
	t := Triangle{base: 3, height: 1}
	fmt.Println("reciever is a value ", t.area())

	//no asterisk in changebase - reciver is value
	//works on a copy of triangle struct so can't change initised struct

	tt := Triangle{base: 3, height: 1}
	tt.changeBase(4)
	fmt.Println("changebase using only value, no change", tt.base)

	tr := Triangle{base: 3, height: 1}
	tr.changeRefBase(4)
	fmt.Println("changebase passing by ref so changes it", tr.base)

}
