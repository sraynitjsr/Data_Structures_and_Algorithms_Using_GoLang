// Final Commit Done

package tricky

import "fmt"

func trickyCodeAppend(data []int) {
	data[2] = -1
	fmt.Println("2 => ", data)
	data = append(data, 2002) // This changes refernce now for data
	fmt.Println("3 => ", data)
	data[3] = 8080
	fmt.Println("4 => ", data)
}

func TrickySlice2() {
	fmt.Println("Tricky SLice Usage 2")
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("1 => ", data)
	trickyCodeAppend(data)
	fmt.Println("5 => ", data)
}
