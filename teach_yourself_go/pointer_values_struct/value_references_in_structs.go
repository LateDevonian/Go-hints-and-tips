package main

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := a
	b.Ice = false
	fmt.Println("assign a directly to be copies a sep version of a")
	fmt.Printf("The value of a %+v\n", a)
	fmt.Printf("The value of b %+v\n", b)
	fmt.Printf("The memory address of a %p\n", &a)
	fmt.Printf("The memory address of b %p\n", &b)

	c := &a
	fmt.Println("assign &a to c - a value by reference")
	fmt.Printf("The value of a, copied as a ref to c %+v\n", a)
	fmt.Printf("The value of c which is a pointer to &a %+v\n", *c)
	fmt.Printf("The underlying memory address of &a %p\n", &a)
	fmt.Printf("The underlying memory address of c %p\n", c)

}
