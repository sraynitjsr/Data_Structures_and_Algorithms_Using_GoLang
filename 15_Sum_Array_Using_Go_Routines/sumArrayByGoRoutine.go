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

	slicesChan := make(chan []int, 10)
	resultsChan := make(chan int, 10)

	go worker(slicesChan, resultsChan)

	for _, newSlice := range myInputSlice {
		slicesChan <- newSlice
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-resultsChan)
	}

	close(slicesChan)
	close(resultsChan)
}
