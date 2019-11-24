package main

import (
	"fmt"
)

type Movie struct {
	Name   string
	Rating float32
}

type Album struct {
	Name    string
	Artist  string
	Trackno float32
}

func main() {
	//can also var m Movie to initalise it
	m := Movie{
		Name:   "Heathers",
		Rating: 8,
	}
	var al Album

	fmt.Println("prints initiatlised by struct contents", m.Name, m.Rating)
	//to print out field names and values of a struct
	fmt.Println("print feild namesa and values %+v\n", m)
	//will print zero values if not initalised
	fmt.Println("print nothing for uninitialsied %+v\n", al)

	//can assign values using dot notation
	var fifth Movie
	fifth.Name = "Fifth Element"
	fifth.Rating = 100

	fmt.Println("prints after initalising indivicually", fifth.Name, fifth.Rating)

	m2 := new(Movie)
	m2.Name = "Milk"
	m2.Rating = 1
	fmt.Println("initalize using new(Movie)", m2.Name, m2.Rating)

	m3 := Movie{Name: "Cats",
		Rating: 2}
	fmt.Println("initalize using Movie(initalise)", m3.Name, m3.Rating)
}
