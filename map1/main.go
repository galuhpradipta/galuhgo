package main

import (
	"fmt"
)

func main() {

	var mapPerson map[int]string

	mapPerson = make(map[int]string)

	mapPerson[1] = "Alex"
	mapPerson[2] = "Kane"
	mapPerson[3] = "Robb"

	for k, v := range mapPerson {
		fmt.Println(k, v)
	}

	galuh, ok := mapPerson[4]
	if !ok {
		fmt.Println("Galuh is not on mapPerson")
	} else {
		fmt.Println(galuh)
	}
}
