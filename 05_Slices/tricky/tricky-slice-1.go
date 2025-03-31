// Final Commit Done

package tricky

import "fmt"

func trickyCode(data []int) {
	data[2] = 20
	fmt.Println("Two ", data)
}

func TrickySlice1() {
	fmt.Println("Tricky SLice Usage 1")
	data := []int{1, 3, 5, 7, 9}
	fmt.Println("One ", data)
	trickyCode(data)
	fmt.Println("Three", data)
}
