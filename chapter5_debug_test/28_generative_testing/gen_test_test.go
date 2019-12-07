package main

import (
	"log"
	"strings"
	"testing"
	"testing/quick"
)

//set up, functiont hat pads a string to a lenghth or truncates
//it if its greater than length

func Pad(s string, max uint) string {
	//logs output
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	ln := uint(len(s))
	if ln > max {
		//if string is loonger than max, truncate it

		return s[:max-1]

	}
	//pad it till correct length
	s += strings.Repeat("", int(max-ln))
	return s
}

//test
func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		//fn takes a string and a uint8, runs Pad()
		//and checks returned length is right
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}
	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		//using testing quick tell it to run no more than 200 randomly
		//generated tests of fn
		t.Error(err)
		//report errors throuough normal patterns
	}
}
