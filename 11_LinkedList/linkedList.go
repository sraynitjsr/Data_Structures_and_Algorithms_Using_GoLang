package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	start *node
}

func (l *list) insertAtBeginning(item int) {
	newNode := &node{
		data: item,
		next: nil,
	}

	temp := l.start

	if temp == nil {
		l.start = newNode
		return
	}

	newNode.next = temp
	l.start = newNode
}

func (l *list) deleteAtBeginning() {
	temp := l.start

	if temp == nil {
		fmt.Println("Nothing present to be deleted")
		return
	}

	l.start = l.start.next
	fmt.Printf("Data deleted from the list => %v", temp.data)

	fmt.Println()
}

func (l *list) display() {
	temp := l.start

	if temp == nil {
		fmt.Println("Empty List")
		return
	}

	fmt.Print("Current List is => ")
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}

	fmt.Println()
}

func Start() {
	fmt.Println("Inside LinkedList")

	myList := &list{}
	myList.insertAtBeginning(10)
	myList.insertAtBeginning(30)
	myList.insertAtBeginning(20)
	myList.insertAtBeginning(50)
	myList.insertAtBeginning(40)

	myList.display()

	myList.deleteAtBeginning()

	myList.deleteAtBeginning()

	myList.display()
}

