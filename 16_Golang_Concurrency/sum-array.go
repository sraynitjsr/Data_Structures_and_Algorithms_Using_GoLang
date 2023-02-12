package golangconcurrency

import "fmt"

func firstSum(firstSlice []int, firstChan chan int) {
	sum := 0
	for _, data := range firstSlice {
		sum = sum + data
	}
	firstChan <- sum
	close(firstChan)
}

func secondSum(secondSlice []int, secondChan chan int) {
	sum := 0
	for _, data := range secondSlice {
		sum = sum + data
	}
	secondChan <- sum
	close(secondChan)
}

func sumArray() {
	fmt.Println("Inside Summ Array Elements From Two Routines Module")
	mySlice := []int{1, 2, 3, 4, 5, 6}
	firstPart := mySlice[:len(mySlice)/2]
	lastPart := mySlice[len(mySlice)/2:]
	myFirstChannel := make(chan int, 1)
	mySecondChannel := make(chan int, 1)
	go firstSum(firstPart, myFirstChannel)
	go secondSum(lastPart, mySecondChannel)
	fmt.Println("First Part Sum =>", <-myFirstChannel)
	fmt.Println("Second Part Sum =>", <-mySecondChannel)
}
