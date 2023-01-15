package hashmap

import "fmt"

func Start() {
	fmt.Println("Inside HashMap")

	myMap := make(map[string]int)
	myMap["A"] = 10
	myMap["C"] = 30
	myMap["B"] = 20
	for key, value := range myMap {
		fmt.Print(key, " ", value)
	}
	keyToBeChecked := "B"
	valueAtKey, keyPresentOrNot := myMap[keyToBeChecked]
	if keyPresentOrNot {
		fmt.Println("\nFound", valueAtKey, "at", keyToBeChecked)
	} else {
		fmt.Println("\nNot Found", keyPresentOrNot)
	}

	delete(myMap, "C")
	_, dataAtCPresentOrNot := myMap["C"]
	if !dataAtCPresentOrNot {
		fmt.Println("No data found at C")
	}
}
