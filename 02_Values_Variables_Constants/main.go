package main

import (
	"fmt"
)

var myGlobalString = "My Global String"

const myConstantVariableOne string = "I am inevitable"

func main() {
	fmt.Println("String" + " Plus " + "String")
	fmt.Println("Int Plus Int =", 1+2)
	fmt.Println("10 + 20 =", 10+20)
	fmt.Println("7.0 / 2 =", 7.0/2)
	fmt.Println(true && false)
	fmt.Println(!false)

	fmt.Println(myGlobalString)
	var myVarString = "My Temp String"
	fmt.Println(myVarString)
	myShortHandString := "My Shorthand String"
	fmt.Println(myShortHandString)
	
	const myConstantVariableTwo string = "And I.... am Iron Man"
	fmt.Println(myConstantVariableOne, myConstantVariableTwo)
}
