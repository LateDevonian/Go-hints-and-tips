package main

import (
	"fmt"
)

//use slices because you can't add elements to an array. slices have
//flexibilty to add, copy, and remove elements, not like an array

func main() {
	//delcare an empty slice with a lenght of 2 slice

	var cheeses = make([]string, 2)
	//the rest works like an array:
	cheeses[0] = "edam"
	cheeses[1] = "blue"
	fmt.Println("basic", cheeses)
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])

	//add elements to a slice
	cheeses = append(cheeses, "gouda")
	fmt.Println("now with append", cheeses)
	fmt.Println(cheeses[2])

	//append is a variadic function
	cheeses = append(cheeses, "havarti", "camembert", "brie")
	fmt.Println("now with even more append", cheeses)

	//delete froma  slice
	//ok this is odd
	fmt.Println("removing ze cheeses first count", len(cheeses))
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println("now with two removed", cheeses)
	fmt.Println("second resize count", len(cheeses))

	//to copy elements
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	//note this only coies the original not appended cheeses. huh
	fmt.Println("copy cheeses to smellyCheeses, print sc ", smellyCheeses)
	//copy individal line:  copy(smellyCheeses, cheeses[1:])
}
