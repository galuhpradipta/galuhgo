package main

import (
	"fmt"
)

// Channel
func main() {

	done := make(chan bool)

	go helloGo(done)

	<-done
	fmt.Println("this is function main")

}

func helloGo(done chan bool) {
	fmt.Println("Hello Golang")
	done <- true
}
