package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//run this with strings ie go run minimize_nils.go hello world

func main() {
	//uyou don't need the if statement, you just use the returned value
	//and y ou've already set up errors.New
	args := os.Args[1:]
	result, _ := Concat(args...)
	fmt.Printf("Concatenated string : '%s'\n", result)

}

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")

	}
	return strings.Join(parts, " "), nil
}
