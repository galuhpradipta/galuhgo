package main

import (
	"fmt"
)

func main() {

	slice := []string{"Back", "End", "Engineer", "DevOps"}

	slice = append(slice, "Techinasia")

	for _, v := range slice {
		fmt.Println(v)
	}
}
