package producer_consumer_async

import "fmt"

func producer(dataChannel chan int, number int) {
	for i := 1; i <= number; i++ {
		fmt.Println("Produced:", i)
		dataChannel <- i
	}
	close(dataChannel)
}

func consumer(dataChannel chan int, waitChannel chan bool) {
	for data := range dataChannel {
		fmt.Println("Consumed:", data)
	}
	waitChannel <- true
}

func Start() {
	fmt.Println("Inside Asynchronized Producer Consumer")
	number := 10
	dataChannel := make(chan int, number)
	waitChannel := make(chan bool)
	go producer(dataChannel, number)
	go consumer(dataChannel, waitChannel)
	<-waitChannel
}
