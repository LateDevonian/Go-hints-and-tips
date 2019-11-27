package main

import (
	"fmt"
)

//map is unorderd group accessed by key rather than an array which is ordered
//this is a hash in ruby
//key values

func main() {
	var players = make(map[string]int)
	players["jean"] = 32
	players["mcangus"] = 14
	players["root"] = 2
	players["coad"] = 8
	fmt.Println("print ALL", players)
	fmt.Println("print by my surname", players["jean"])

	//delete
	delete(players, "coad")
	fmt.Println("print ALL AFTER DELETE", players)

}
