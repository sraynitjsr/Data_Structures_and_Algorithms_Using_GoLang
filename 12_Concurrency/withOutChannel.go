package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func print(wg *sync.WaitGroup, str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Microsecond * 200)
	}
	wg.Done()
}

func StartWithout() {
	var wg sync.WaitGroup
	wg.Add(1)

	print(&wg, "Hello World")

	wg.Wait()
}
