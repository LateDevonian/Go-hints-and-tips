//write a function that takes generic values (interface{}s
//and does somthing useful based on underlying types)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 2
	var b int = 37
	var c string = "3.2"
	result := sum(a, b, c)
	fmt.Printf("result: %f\n", result)
}

//for each type supported, convert to a float 64 and sum
func sum(v ...interface{}) float64 {
	var result float64 = 0
	for _, val := range v {
		switch val.(type) {
		case int:
			result += float64(val.(int))
		case int64:
			result += float64(val.(int64))
		case uint8:
			result += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			result += a
		default:
			fmt.Printf("unsuppoted type %T. Ignoring.\n", val)
		}
	}
	return result
}
