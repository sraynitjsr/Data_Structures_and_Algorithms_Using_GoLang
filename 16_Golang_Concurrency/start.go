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
}

