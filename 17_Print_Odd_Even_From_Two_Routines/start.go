package printoddevenfromtworoutines

import (
	"fmt"
	"sync"
)

func print(i int, wg *sync.WaitGroup, str string) {
	fmt.Println(str, i)
	defer wg.Done()
}

func Start() {
	fmt.Println("Print Odd Even From Alternate Go Routines")
	var wg sync.WaitGroup
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			wg.Add(1)
			go print(i, &wg, "Even")
		} else {
			wg.Add(1)
			go print(i, &wg, "Odd")
		}
		wg.Wait()
	}
}
