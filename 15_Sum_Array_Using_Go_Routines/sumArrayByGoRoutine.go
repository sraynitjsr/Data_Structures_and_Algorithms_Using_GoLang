package sumarrayusinggoroutines

import "fmt"

func addSlice(mySlice []int) int {
	sum := 0
	for _, val := range mySlice {
		sum = sum + val
	}
	return sum
}

func worker(slices chan []int, results chan int) {
	for newSlicew := range slices {
		results <- addSlice(newSlicew)
	}
}

func Start() {
	myInputSlice := [][]int{}
	myInputSlice = append(myInputSlice, []int{10, 20, 30})
	myInputSlice = append(myInputSlice, []int{40, 50, 60})
	myInputSlice = append(myInputSlice, []int{70, 80, 90})

	slicesChan := make(chan []int, 3)
	resultsChan := make(chan int, 3)

	for _, newSlice := range myInputSlice {
		slicesChan <- newSlice
	}

	go worker(slicesChan, resultsChan)

	for i := 0; i < 3; i++ {
		fmt.Println("Current Slice Sum => ", <-resultsChan)
	}

	close(slicesChan)
	close(resultsChan)
}
