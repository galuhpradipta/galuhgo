package main

import (
	"fmt"
)

func main() {

	var players []Player

	players = []Player{Player{ID: 1, Name: "Cristiano Ronaldo"}, Player{ID: 2, Name: "Lionel Messi"}}

	players = append(players, Player{ID: 3, Name: "Harry Kane"})

	for _, v := range players {
		fmt.Println(v)
	}
}

type Player struct {
	ID   int
	Name string
}
