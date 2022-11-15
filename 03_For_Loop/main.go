package main

import (
	"fmt"
)

func main() {
	myItr1 := 0
	for myItr1 < 5 {
		fmt.Print(myItr1, " ")
		myItr1++
	}

	fmt.Println()

	for myItr2 := 10; myItr2 < 15; myItr2++ {
		fmt.Print(myItr2, " ")
	}

	fmt.Println()

	for {
		fmt.Println("Infinite Loop, Press Force")
	}
}
