package concurrency

import "fmt"

func addToChannel(ch chan string, data string) {
	ch <- data
}

func readFromChannel(ch chan string) {
	count := 0
	for count < 5 {
		fmt.Println(<-ch)
		count++
	}
}

func StartLater() {
	fmt.Println("Inside Concurrency")

	myChannel := make(chan string)
	go addToChannel(myChannel, "A")
	go addToChannel(myChannel, "C")
	go addToChannel(myChannel, "B")
	go addToChannel(myChannel, "E")
	go addToChannel(myChannel, "D")
	readFromChannel(myChannel)
}
