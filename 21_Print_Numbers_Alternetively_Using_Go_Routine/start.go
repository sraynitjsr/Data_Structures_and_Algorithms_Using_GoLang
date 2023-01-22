package printnumbersalternetivelyusinggoroutine

import (
	"fmt"
	"time"
)

func method1(ch chan int) {
	for {
		ch <- 100
	}
}

func method2(ch chan int) {
	for {
		ch <- 200
	}
}

func Start() {
	fmt.Println("Inside For Select Example")
	myChannel := make(chan int)

	go method1(myChannel)
	go method2(myChannel)

	time.Sleep(time.Second * 2)

	for {
		select {
		case msg1 := <-myChannel:
			time.Sleep(time.Second * 1)
			fmt.Println(msg1)
		case msg2 := <-myChannel:
			fmt.Println(msg2)
			time.Sleep(time.Second * 1)
		}
	}
}
