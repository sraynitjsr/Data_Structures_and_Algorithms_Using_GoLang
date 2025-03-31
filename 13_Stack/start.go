// Final Commit Done

package stack

import "fmt"

type MyStack struct {
	mySlice []interface{}
}

func (myStack *MyStack) push(val interface{}) {
	myStack.mySlice = append(myStack.mySlice, val)
}

func (myStack *MyStack) pop() {
	if len(myStack.mySlice) == 0 {
		fmt.Println("No Data in the Stack")
		return
	}
	sizeOfStack := len(myStack.mySlice)
	myStack.mySlice = myStack.mySlice[:sizeOfStack-1]
}

func (myStack *MyStack) display() {
	if len(myStack.mySlice) == 0 {
		fmt.Println("No Data in the Stack")
		return
	}
	fmt.Print("Current Data in the Stack is => ")
	sizeOfStack := len(myStack.mySlice)
	for i := sizeOfStack - 1; i >= 0; i-- {
		fmt.Print(myStack.mySlice[i], " ")
	}
	fmt.Println("")
}

func Start() {
	fmt.Println("Stack in GoLang")

	myStack := &MyStack{}

	myStack.pop()

	myStack.push(10)
	myStack.push(20)
	myStack.push(30)
	myStack.push(40)

	myStack.display()

	myStack.pop()
	myStack.pop()

	myStack.display()

	myStack.pop()

	myStack.push(50)

	myStack.display()
}
