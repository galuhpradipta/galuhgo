package main

import (
	"fmt"
)

func main() {

	nextValue := genNumber()

	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	lv := love("Galuh Pradipta")
	fmt.Println(lv("Golang"))
	fmt.Println(lv("Programming"))

}

func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s very love %s", name, things)
	}
}
