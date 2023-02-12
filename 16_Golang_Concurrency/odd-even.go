package golangconcurrency

import (
	"fmt"
	"sync"
)

func sendOdd(num int, wg *sync.WaitGroup, myDoneChannel chan bool) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		if i%2 != 0 && <-myDoneChannel {
			fmt.Println("Odd", i)
			myDoneChannel <- true
		}
	}
}

func sendEven(num int, wg *sync.WaitGroup, myDoneChannel chan bool) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		if i%2 == 0 && <-myDoneChannel {
			fmt.Println("Even", i)
			myDoneChannel <- true
		}
	}
}

func oddEven() {
	fmt.Println("Inside Print Odd Even From Two Routines Alternatively Module")
	var wg sync.WaitGroup
	myNumber := 20
	myDoneChannel := make(chan bool, 1)
	myDoneChannel <- true
	wg.Add(2)
	go sendEven(myNumber, &wg, myDoneChannel)
	go sendOdd(myNumber, &wg, myDoneChannel)
	wg.Wait()
}
