package main

import (
	"fmt"
)

func main() {

	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arrString := [3]string{"Back", "End", "Engineer"}

	for i := 0; i < len(arrString); i++ {
		fmt.Println(arrString[i])
	}

	arrMulti := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(arrMulti[0][2])
}
