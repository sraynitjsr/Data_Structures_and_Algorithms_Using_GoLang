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

func paniDemo() {
	fmt.Println("Starting program...")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	divide(10, 1)

	fmt.Println("Program continues...")

	divide(10, 0)
}

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func Start() {
	fmt.Println("Topics Missed Previously, Runes")
	traverseAString()
	fmt.Println()
	usingRuneInGoLang()
	fmt.Println()
	paniDemo()
}
