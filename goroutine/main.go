package main

import (
	"fmt"
	"time"
)

func main() {

	go helloGo()

	time.Sleep(3 * time.Millisecond)

	fmt.Println("Hello to Golang")
}

func helloGo() {
	fmt.Println("This is function")
}
