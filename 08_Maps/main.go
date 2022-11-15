package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["Sachin"] = 100
	myMap["Sourav"] = 200
	myMap["Virat"] = 300
	myMap["Lara"] = 400
	fmt.Println(myMap)

	valueAtKeySachin, valevalueAtKeySachinYesOrNo := myMap["Lara"]
	fmt.Println(valueAtKeySachin, valevalueAtKeySachinYesOrNo)

	delete(myMap, "Lara")

	valueAtKeyLara, valevalueAtKeyLaraYesOrNo := myMap["Lara"]
	fmt.Println(valueAtKeyLara, valevalueAtKeyLaraYesOrNo)
}
