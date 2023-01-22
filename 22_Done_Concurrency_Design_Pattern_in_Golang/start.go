package done_design_pattern

import (
	"fmt"
	"time"
)

func doSomeTask(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Hello World")
			time.Sleep(time.Second * 3)
		}
	}
}

func Start() {
	done := make(chan bool)
	go doSomeTask(done)
	time.Sleep(time.Second * 10)
}
