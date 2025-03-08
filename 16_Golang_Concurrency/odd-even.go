package golangconcurrency

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	writeWaitGroup sync.WaitGroup

	lock      sync.Mutex
	condition = sync.NewCond(&lock)

	myNumber  = 10
	flag      = true
	myChannel = make(chan string)
)

func oddWriter() {
	defer writeWaitGroup.Done()
	for i := 1; i <= myNumber; i += 2 {
		lock.Lock()
		for !flag {
			condition.Wait()
		}
		myChannel <- "Odd  Data => " + strconv.Itoa(i)
		flag = false
		condition.Signal()
		lock.Unlock()
	}
}

func evenWriter() {
	defer writeWaitGroup.Done()
	for i := 2; i <= myNumber; i += 2 {
		lock.Lock()
		for flag {
			condition.Wait()
		}
		myChannel <- "Even Data => " + strconv.Itoa(i)
		flag = true
		condition.Signal()
		lock.Unlock()
	}
}

func oddEven() {
	fmt.Println("Inside Print Odd Even From Two Routines Alternatively Module")

	writeWaitGroup.Add(2)
	go oddWriter()
	go evenWriter()

	go func() {
		writeWaitGroup.Wait()
		close(myChannel)
	}()

	for data := range myChannel {
		fmt.Println(data)
	}
}
