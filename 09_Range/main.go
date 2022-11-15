package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 3)
	mySlice[0] = 100
	mySlice[1] = 101
	mySlice[2] = 102
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	for index, value := range mySlice {
		fmt.Println(index, value)
	}
}
