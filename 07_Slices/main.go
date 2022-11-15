package main

import (
	"fmt"
)

func main() {
	mySlice := make([]string, 5)
	mySlice[0] = "A"
	mySlice[1] = "B"
	mySlice[2] = "C"
	mySlice[3] = "D"
	mySlice[4] = "E"
	fmt.Println(mySlice)

	mySlice = append(mySlice, "F")
	mySlice = append(mySlice, "G-to-Z")
	fmt.Println(mySlice)

	fmt.Println(mySlice[0 : len(mySlice)-2])

	twoDimensionalIntegerSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerSlice := make([]int, 3)
		for j := 0; j < 3; j++ {
			innerSlice[j] = i + j + 1
		}
		twoDimensionalIntegerSlice[i] = innerSlice
	}

	for i := 0; i < 3; i++ {
		fmt.Println(twoDimensionalIntegerSlice[i])
	}
}
