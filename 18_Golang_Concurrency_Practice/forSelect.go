package golangconcurrencypractice

import "fmt"

func method1(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "From Channel One"
	}
}

func method2(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "From Channel Two"
	}
}

func forSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go method1(ch1)
	go method2(ch2)

	select {
	case c1 := <-ch1:
		fmt.Println(c1)
	case c2 := <-ch2:
		fmt.Println(c2)
	}
}
