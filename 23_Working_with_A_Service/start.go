package workingwithaservice

import (
	"fmt"
	"time"
)

func doTask(ch chan string, str string) {
	for {
		ch <- str
	}
}

func Start() {
	fmt.Println("Inside Running Service Module")
	myChannel := make(chan string)
	go doTask(myChannel, "Hello World")
	for data := range myChannel {
		fmt.Println("Current Data From Channel is =>", data)
		time.Sleep(time.Second * 1)
	}
}
