package golangconcurrency

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	numberOfTasks    = 15
	numberOfWorkers  = 3
	workersWaitGroup sync.WaitGroup
	tasksChannel     = make(chan int, numberOfTasks)
	resultsChannel   = make(chan string, numberOfTasks)
)

func workerLogic(workerID int) {
	defer workersWaitGroup.Done()
	for task := range tasksChannel {
		time.Sleep(time.Second)
		resultsChannel <- "Worker Number " + strconv.Itoa(workerID) + " Doing Task Number " + strconv.Itoa(task)
	}
}

func createWorkers() {
	for w := 1; w <= numberOfWorkers; w++ {
		workersWaitGroup.Add(1)
		go workerLogic(w)
	}
}

func createTasksAndCloseTasksChannel() {
	for t := 1; t <= numberOfTasks; t++ {
		tasksChannel <- t
	}
	close(tasksChannel)
}

func waitForWorkersToEndAndCloseResultsChannelAndRead() {
	workersWaitGroup.Wait()
	close(resultsChannel)

	for data := range resultsChannel {
		fmt.Println(data)
	}
}

func workerPool() {
	fmt.Println("GoLang Worker Pool Concept")
	createWorkers()
	createTasksAndCloseTasksChannel()
	waitForWorkersToEndAndCloseResultsChannelAndRead()
}
