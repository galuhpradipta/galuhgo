package main

import (
	"fmt"
)

type Vector struct {
	X, Y int
}

type Player struct {
	ID   int
	name string
}

func main() {

	var v Vector
	v.X = 10
	v.Y = 15

	fmt.Println(v)

	fmt.Println("X is = ", v.X)
	fmt.Println("Y is = ", v.Y)

	Player1 := Player{ID: 10, name: "Messi"}

	fmt.Println(Player1.ID)
	fmt.Println(Player1.name)
}
