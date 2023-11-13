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
	queue "github.com/sraynitjsr/12_Queue"
	stack "github.com/sraynitjsr/13_Stack"
	func_prog "github.com/sraynitjsr/19_Functional_Programming"
	jsonmarshalunmarshal "github.com/sraynitjsr/14_JSON_Marshal_UnMarshal"
	bst "github.com/sraynitjsr/15_Binary_Search_Tree"
	concurrency "github.com/sraynitjsr/16_Golang_Concurrency"
	restapi "github.com/sraynitjsr/17_Simple_REST_API"
	svc "github.com/sraynitjsr/18_Working_With_A_Service"
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
	queue.Start()
	jsonmarshalunmarshal.Start()
	stack.Start()
	func_prog.Start()
	bst.Start()
	restapi.Start()
	svc.Start()
	concurrency.Start()
}
