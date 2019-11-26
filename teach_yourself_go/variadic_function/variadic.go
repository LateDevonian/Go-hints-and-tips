package main

import (
	"fmt"
)

//variadic functions accept a  variable number of arguements
//can pass varying number of the same type with ... in the arguements
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumNumbers(1, 2, 3, 4)
	fmt.Println("the result is\n", result)
}
