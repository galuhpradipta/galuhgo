package main

import (
	"fmt"
)

// Channel buffer

// unbuffer channel
// done <- true (sender)
// <- (receiver)

func main() {
	hello := make(chan string, 2)

	hello <- "Hello"
	hello <- "Golang"
	close(hello)

	for v := range hello {
		fmt.Println(v)
	}
}
