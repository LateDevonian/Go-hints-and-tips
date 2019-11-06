package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

//reads yaml from a string or file

func main() {
	config, err := yaml.ReadFile("config.yaml")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
