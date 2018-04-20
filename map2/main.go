package main

import (
	"fmt"
)

func main() {

	mapPlayer := make(map[int]Player)
	mapPlayer[1] = Player{ID: 1, Name: "Torres"}
	mapPlayer[2] = Player{ID: 2, Name: "Silva"}
	mapPlayer[3] = Player{ID: 3, Name: "Kane"}

	for _, v := range mapPlayer {
		fmt.Println(v.ID, v.Name)
	}

}

type Player struct {
	ID   int
	Name string
}
