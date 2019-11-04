package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

//create a struct for the flag options
var opts struct {
	Name    string `short:"n" long:"name" default:"world" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish"`
}

//original code, import "flag"
// var name = flag.String("name", "World", "A name to say hello to")
// var spanish bool

// func init() {
// 	flag.BoolVar(&spanish, "spanish", true, "Use Spanish language.")
// 	flag.BoolVar(&spanish, "s", true, "Use Spanish Language.")
// }

func main() {
	flags.Parse(&opts)

	if opts.Spanish == false {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
