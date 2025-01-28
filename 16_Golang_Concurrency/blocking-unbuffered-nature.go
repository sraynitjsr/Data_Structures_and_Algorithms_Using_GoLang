package golangconcurrency

import "fmt"

func blockingNatureUnbuffered() {
	ch := make(chan int)

	go func() {
		fmt.Println("Sending data to channel...")
		ch <- 8080
		fmt.Println("Data sent!")
	}()

	fmt.Println("Receiving data from channel...")
	data := <-ch
	fmt.Printf("Received: %d\n", data)
}
