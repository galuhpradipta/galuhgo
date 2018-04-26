package main

import (
	"fmt"
)

func main() {
	fmt.Println("Worker Pool")
}

func FakeHttpRequest(x int) {
	return x * 10
}

func producer(jobs chan<- int, size int) {
	for i := 1; i < size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int) {

}
