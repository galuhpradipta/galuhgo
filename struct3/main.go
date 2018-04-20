package main

import (
	"fmt"
)

func main() {

	p := &Person{
		ID:   1,
		Name: "Galuh",
		Age:  23,
	}

	fmt.Println(p.GetName())
	p.changeName("Galuh Pradipta")
	fmt.Println(p.GetName())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}
