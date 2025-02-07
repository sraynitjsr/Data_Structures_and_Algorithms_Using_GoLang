package missedtopics

import "fmt"

func traverseAString() {
	myStr := "Some_Dummy_String"
	for _, v := range myStr {
		if string(v) == "_" {
			fmt.Println()
		} else {
			fmt.Print(string(v))
		}
	}
}

func usingRuneInGoLang() {
	mySlice := []rune{}
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, rune(i))
	}
	fmt.Println(mySlice)
}

func Start() {
	fmt.Println("Topics Missed Previously, Runes")
	traverseAString()
	fmt.Println()
	usingRuneInGoLang()
}
