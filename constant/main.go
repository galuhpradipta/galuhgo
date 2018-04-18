package main

import (
	"fmt"
)

const A string = "ini konstan"

func main() {
	fmt.Println(A)

	const I int = 10

	fmt.Println(I)

	x := 500 / I

	fmt.Println(x)
}
