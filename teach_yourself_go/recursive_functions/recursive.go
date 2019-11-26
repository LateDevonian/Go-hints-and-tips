package main

import (
	"fmt"
)

//this is how you do recursioin in go!

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d/n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! i've eaten %d/n", eaten)
	return feedMe(portion, eaten)
}

func main() {
	fmt.Println(feedMe(1, 0))
}
