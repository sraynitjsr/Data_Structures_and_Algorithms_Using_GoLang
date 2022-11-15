package main

import (
	"fmt"
)

func main() {

	var oneDimensionalStringArray [3]string
	oneDimensionalStringArray[0] = "A"
	oneDimensionalStringArray[1] = "B"
	oneDimensionalStringArray[2] = "C"

	for i := 0; i < len(oneDimensionalStringArray); i++ {
		fmt.Print(oneDimensionalStringArray[i], " ")
	}

	fmt.Println()

	oneDimensionalIntArray := [3]int{1, 2, 3}
	for j := 0; j < len(oneDimensionalIntArray); j++ {
		fmt.Print(oneDimensionalIntArray[j], " ")
	}

	fmt.Println()

	var twoDimensionalIntArray [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoDimensionalIntArray[i][j] = i + j
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(twoDimensionalIntArray[i][j], " ")
		}
		fmt.Println()
	}
}
