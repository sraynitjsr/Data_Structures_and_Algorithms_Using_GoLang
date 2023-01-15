package basicsyntaxgo

import (
	"fmt"
	"reflect"
)

const x int = 20

func Start() {
	fmt.Println("Inside Basic Syntax of Go")

	fmt.Println("X is", x)

	var b string = "B"
	fmt.Println(b)

	c := 35.46
	fmt.Println("Value of c is", c, "of type", reflect.TypeOf(c).Name())

	data := 100
	fmt.Println(data)
	data++
	fmt.Println(data)
}
