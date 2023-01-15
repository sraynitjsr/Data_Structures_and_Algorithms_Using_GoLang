package main

import (
	"fmt"

	helloworld "github.com/sraynitjsr/01_Hello_World"
	basics "github.com/sraynitjsr/02_Basic_Syntax_Go"
	io "github.com/sraynitjsr/03_Input_Output"
	arrays "github.com/sraynitjsr/04_Arrays"
	slices "github.com/sraynitjsr/05_Slices"
	sorting "github.com/sraynitjsr/06_Sorting"
	hashmap "github.com/sraynitjsr/07_HashMap"
	interfaces "github.com/sraynitjsr/08_Interface"
	hashset "github.com/sraynitjsr/09_HashSet"
	mystrings "github.com/sraynitjsr/10_Strings"
	linkedlist "github.com/sraynitjsr/11_LinkedList"
	concurrency "github.com/sraynitjsr/12_Concurrency"
	queue "github.com/sraynitjsr/13_Queue"
	jsonmarshalunmarshal "github.com/sraynitjsr/14_JSON_Marshal_UnMarshal"
	sumbyroutine "github.com/sraynitjsr/15_Sum_Array_Using_Go_Routines"
	restapi "github.com/sraynitjsr/16_Simple_REST_API"
)

func main() {
	fmt.Println("Inside Main")
	helloworld.Start()
	basics.Start()
	io.Start()
	arrays.Start()
	slices.Start()
	sorting.Start()
	hashmap.Start()
	interfaces.Start()
	hashset.Start()
	hashset.Start_Optimized()
	mystrings.Start()
	linkedlist.Start()
	concurrency.Start()
	queue.Start()
	jsonmarshalunmarshal.Start()
	sumbyroutine.Start()
	restapi.Start() // Need to close the program, otherwise the control won't move forward to next parts, as the application is always running and listening to the given port
}
