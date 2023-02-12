package queue

import "fmt"

type Queue struct {
	items []*interface{}
}

func (q *Queue) enqueue(item interface{}) {
	q.items = append(q.items, &item)
}

func (q *Queue) dequeue() {
	if len(q.items) == 0 {
		return
	}

	first := q.items[0]
	fmt.Println(*first, "deleted from the queue")

	q.items = q.items[1:]
}

func (q *Queue) display() {
	if len(q.items) == 0 {
		fmt.Println("Empty Quque")
		return
	}

	fmt.Print("Current Queue => ")
	for i := 0; i < len(q.items); i++ {
		fmt.Print("<", *q.items[i], "> ")
	}
	fmt.Println()
}

func Start() {
	fmt.Println("Inside Queue")

	myQueue := &Queue{}

	myQueue.display()

	myQueue.enqueue("Sachin")
	myQueue.enqueue("Sourav")
	myQueue.enqueue("Virat")

	myQueue.display()

	myQueue.dequeue()

	myQueue.display()
}

