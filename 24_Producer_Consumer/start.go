package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		ch <- n
		fmt.Println("Produced:", n)
		time.Sleep(time.Second)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("Consumed:", n)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	close(ch)
}
