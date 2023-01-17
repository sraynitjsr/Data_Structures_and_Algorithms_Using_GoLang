package printoddevenfromtworoutines

import (
	"fmt"
	"sync"
)

func Start() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println(i)
			}(i)
			wg.Wait()
		} else {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println(i)
			}(i)
			wg.Wait()
		}
	}
}
