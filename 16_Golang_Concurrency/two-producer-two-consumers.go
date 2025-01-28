package golangconcurrency

import (
	"fmt"
	"sync"
)

func produceToChannel(myChannel chan int, wg1 *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		myChannel <- i
	}
	defer wg1.Done()
}

func consumeFromChannel(myChannel chan int, wg2 *sync.WaitGroup) {
	for data := range myChannel {
		fmt.Println("Consuming Now =>", data)
	}
	defer wg2.Done()
}

func twoProducerTwoConsumers() {
	myChannel := make(chan int)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	count := 2
	for i := 0; i < count; i++ {
		wg1.Add(1)
		go produceToChannel(myChannel, &wg1)
	}
	for i := 0; i < count; i++ {
		wg2.Add(1)
		go consumeFromChannel(myChannel, &wg2)
	}
	wg1.Wait()
	close(myChannel)
	wg2.Wait()
}
