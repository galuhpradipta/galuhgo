package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	// Error handling when we try to conver string to integer
	iStr := "A"

	i, err := strconv.Atoi(iStr)

	if err != nil {
		fmt.Println("Error happens", err.Error())
	} else {
		fmt.Println(i)
	}

	r, err := Div(10, 0)

	if err != nil {
		fmt.Println("Error happens", err.Error())
	} else {
		fmt.Println(r)
	}

	// inline error handling

	if x, err := Div(25, 5); err != nil {
		fmt.Println("Error happen", err.Error())
	} else {
		fmt.Println(x)
	}
}

func Div(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("Cant divide with zero")
	} else {
		return x / y, nil
	}
}
