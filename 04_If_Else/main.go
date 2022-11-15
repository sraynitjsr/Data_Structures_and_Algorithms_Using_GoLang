package main

import (
	"fmt"
)

func main() {
	if 1 > 2 {
		fmt.Println("1 > 2")
	} else if 1 < 2 {
		fmt.Println("1 < 2")
	} else {
		fmt.Println("1 == 2")
	}

	if tempVar := 13; tempVar%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
