package arrays

import (
	"fmt"
	"strconv"
)

func firstArray(arr [5]int) {
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	for index, value := range arr {
		fmt.Printf("Value at index %v is %v\n", index, value)
	}
}

func secondArray(myArray [5]string) {
	for i := 0; i < 5; i++ {
		myArray[i] = "String => " + strconv.Itoa(i)
	}
	for in, val := range myArray {
		fmt.Printf("Value present at index %v is %v\n", in, val)
	}
}

func Start() {
	fmt.Println("Inside Arrays")
	my1DArray := [5]int{}
	firstArray(my1DArray)
	my2DArray := [5]string{}
	secondArray(my2DArray)
}

