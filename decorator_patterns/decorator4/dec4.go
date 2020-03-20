//https://stackoverflow.com/questions/45944781/decorator-functions-in-go
package main

import (
	"fmt"
	"strings"
)

const (
	prefix = "PREFIX"
	suffix = "SUFFIX"
)

//uses golangs type aliases, decorated and decorator
type Decorated = string

func addConstPrefix(s string) string {
	return prefix + s
}

func addConstSuffix(s string) string {
	return s + suffix
}

//a parametized decorator is a function which returns another function

func addPrefix(prefix string) func(string) string {
	return func(s string) string {
		return prefix + s
	}
}

type Decorator = func(string) string

func Decorate(c Decorated, ds ...Decorator) Decorated {
	decorated := c
	for _, decorator := range ds {
		decorated = decorator(decorated)
	}
	return decorated
}

func main() {
	//creating a few decorators (functions) which match the type sig
	//we defined

	var toLower Decorator = strings.ToLower
	var toUpper Decorator = strings.ToUpper
	var dec3 Decorator = addConstPrefix
	var dec4 Decorator = addConstSuffix

	input := "Let's decorate me!"
	inputUppercase := Decorate(input, []Decorator{toUpper}...)
	fmt.Println(inputUppercase)

	//apply the decorators

	allDecorators := []Decorator{toUpper, toLower, dec3, dec4}
	output := Decorate(input, allDecorators...)
	fmt.Println(output)

	input2 := "Let's decorate me!"
	prefixed := Decorate(input2, []Decorator{addPrefix("Well, ")}...)
	fmt.Println(prefixed)
}
