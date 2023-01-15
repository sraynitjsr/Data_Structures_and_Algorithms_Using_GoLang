package concurrency

import (
	"fmt"
)

func display(ch chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}

func insert(ch chan string, str string) {
	for i := 0; i < 5; i++ {
		ch <- str
	}
	close(ch)
}

func Start() {
	myChannel := make(chan string)
	go insert(myChannel, "Hi Bye")
	display(myChannel)
}
