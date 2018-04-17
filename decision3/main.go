package main

import (
	"fmt"
)

func main() {

	x := 100

	// expression switch
	switch x {
	case 60:
		fmt.Println("SCORE = C")
	case 70:
		fmt.Println("SCORE = B")
	case 90:
		fmt.Println("SCORE = A")
	default:
		fmt.Println("SCORE NOT FOUND")
	}

	switch {
	case x == 60:
		fmt.Println("SCORE = C")
	case x == 70:
		fmt.Println("SCORE = B")
	case x == 90:
		fmt.Println("SCORE = A")
	default:
		fmt.Println("SCORE NOT FOUND")
	}

	// type switch

	var y interface{}
	y = 5

	switch y.(type) {
	case int:
		fmt.Println("y type data is integer")
	case string:
		fmt.Println("y type data is string")
	case float64:
		fmt.Println("y type data is float64")
	default:
		fmt.Println("type data not found")
	}

}
