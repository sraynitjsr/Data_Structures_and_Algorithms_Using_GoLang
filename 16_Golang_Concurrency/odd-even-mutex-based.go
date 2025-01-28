package golangconcurrency

import (
	"fmt"
	"sync"
)

func odd(n int, myChannel chan int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			m.Lock()
			myChannel <- i
			m.Unlock()
		}
	}
}

func even(n int, myChannel chan int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			m.Lock()
			myChannel <- i
			m.Unlock()
		}
	}
}

func oddEvenMutexBased() {
	myChannel := make(chan int)
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(2)
	go odd(10, myChannel, &wg, &m)
	go even(10, myChannel, &wg, &m)

	go func() {
		wg.Wait()
		close(myChannel)
	}()

	for data := range myChannel {
		fmt.Println(data)
	}
}
