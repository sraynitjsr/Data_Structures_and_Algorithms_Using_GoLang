package golangconcurrency

import "fmt"

func producerAsync(dataChannel chan int, number int) {
	for i := 1; i <= number; i++ {
		fmt.Println("Produced:", i)
		dataChannel <- i
	}
	close(dataChannel)
}

func consumerAsync(dataChannel chan int, waitChannel chan bool) {
	for data := range dataChannel {
		fmt.Println("Consumed:", data)
	}
	waitChannel <- true
}

func producerConsumerAsync() {
	fmt.Println("Inside Asynchronized Producer Consumer")
	number := 5
	dataChannel := make(chan int, number)
	waitChannel := make(chan bool)
	go producerAsync(dataChannel, number)
	go consumerAsync(dataChannel, waitChannel)
	<-waitChannel
}

