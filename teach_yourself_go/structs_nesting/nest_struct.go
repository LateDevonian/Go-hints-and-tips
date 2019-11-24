package main

import (
	"fmt"
)

type Superhero struct {
	Name string
	Age  int
	//address is referring to anohter struct
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	e := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain drive",
			City:   "Gotham",
		},
	}
	fmt.Println("what is in the struct: %+v\n", e)
	fmt.Println("access the contents via dot address:", e.Address.Street)

}
