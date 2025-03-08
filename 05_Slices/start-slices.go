package slices

import (
	"fmt"

	"github.com/sraynitjsr/05_Slices/tricky"
)

func basics() {
	my1DIntSlice := []int{}
	for i := 0; i < 3; i++ {
		my1DIntSlice = append(my1DIntSlice, i+1)
	}

	my2DIntSlice := [][]int{}
	for i := 0; i < 3; i++ {
		tempSlice := []int{}
		for j := 0; j < 5; j++ {
			tempSlice = append(tempSlice, 1001)
		}
		my2DIntSlice = append(my2DIntSlice, tempSlice)
	}

	fmt.Println("1D Slice Is")
	for val := range my1DIntSlice {
		fmt.Print(val, " ")
	}

	fmt.Println("\n2D Slice Is")
	for row := range my2DIntSlice {
		fmt.Println(my2DIntSlice[row], " ")
	}
}

func Start() {
	fmt.Println("Inside Slices")

	basics()

	tricky.TrickySlice1()

	tricky.TrickySlice2()
}
