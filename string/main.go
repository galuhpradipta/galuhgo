package main

// string manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var myString = "Hello Golang"
	var myStringTwo = "10"

	fmt.Println(myString)
	fmt.Println(len(myString))

	// Upper

	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)

	// Lower

	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	// Contain

	resultContains := strings.Contains(myString, "Go")
	fmt.Println(resultContains)

	// Split

	resultSplit := strings.Split(myString, " ")
	for _, v := range resultSplit {
		fmt.Println(v)
	}

	// Conversion Str to Int

	resultConvInt, err := strconv.Atoi(myStringTwo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultConvInt)

	// Conversion Int to Str

	resultConvStr := strconv.Itoa(10)
	fmt.Println(resultConvStr)

}
