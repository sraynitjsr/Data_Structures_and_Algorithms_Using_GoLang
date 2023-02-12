package inputoutput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Start() {
	fmt.Println("Inside Input Output")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a string")
	scanner.Scan()
	myString := scanner.Text()

	fmt.Println("Enter an integer")
	scanner.Scan()
	myIntInString := scanner.Text()
	myInt, _ := strconv.ParseInt(myIntInString, 10, 64)

	fmt.Println("Enter a floaing number")
	scanner.Scan()
	myFloatingNumberInString := scanner.Text()
	myFloatingNumber, _ := strconv.ParseFloat(myFloatingNumberInString, 64)

	fmt.Println(myInt, myFloatingNumber, myString)
}

