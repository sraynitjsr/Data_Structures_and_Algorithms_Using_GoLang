package golangconcurrencypractice

import "fmt"

func add(myCurrentChannel chan int, myCurrentSlice []int) {
	sum := 0
	for _, val := range myCurrentSlice {
		sum = sum + val
	}
	myCurrentChannel <- sum
}

func addArrayData() {
	mySlice := []int{1, 2, -3, 4, 5}

	myChannel1 := make(chan int)
	myChannel2 := make(chan int)

	mySliceLen := len(mySlice)

	go add(myChannel1, mySlice[:mySliceLen/2])
	go add(myChannel2, mySlice[mySliceLen/2:])

	fmt.Println(<-myChannel1 + <-myChannel2)

	close(myChannel1)
	close(myChannel2)
}
