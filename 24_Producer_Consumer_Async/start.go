package producer_consumer

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Second)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		n := <-ch
		fmt.Println("Consumed:", n)
	}
}

func Start() {
	fmt.Println("Inside Producer Consumer")
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	close(ch)
}
