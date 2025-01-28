package golangconcurrency

import "fmt"

func Start() {
	fmt.Println("")
	oddEven()
	fmt.Println("")
	sumArray()
	fmt.Println("")
	producerConsumerSync()
	fmt.Println("")
	producerConsumerAsync()
	fmt.Println("")
	twoProducerTwoConsumers()
	fmt.Println("")
	oddEvenMutexBased()
}
