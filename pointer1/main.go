package main

import (
	"fmt"
)

func main() {

	var hello string = "Hello"
	var strPtr *string

	strPtr = &hello

	fmt.Println(hello)
	changePtr(strPtr)
	fmt.Println(hello)

}

// pass by value
func change(v string) {
	v = v + " Golang"
}

// pass by reference
func changePtr(v *string) {
	*v = *v + " Golang"
}
