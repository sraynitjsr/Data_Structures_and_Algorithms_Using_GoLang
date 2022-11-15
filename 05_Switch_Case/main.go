package main

import (
	"fmt"
)

func main() {
	myInt := 10
	switch {
	case myInt < 5:
		fmt.Println("Less")
	case myInt > 5:
		fmt.Println("More")
	default:
		fmt.Println("Equal")
	}
}
