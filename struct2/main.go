package main

import (
	"fmt"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {

	p := Person{
		ID:   1,
		Name: "Galuh",
		Age:  23,
	}

	printPerson(p)
}

func printPerson(p Person) {
	fmt.Println("ID = ", p.ID)
	fmt.Println("Name = ", p.Name)
	fmt.Println("Age = ", p.Age)
}
