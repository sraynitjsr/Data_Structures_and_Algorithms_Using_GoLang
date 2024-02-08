package newstuff

import (
	"fmt"
	"maps"
	"slices"
)

func Start() {
	m1 := map[string]int{
		"Age":  20,
		"Roll": 100,
	}
	m2 := map[string]int{
		"Age":  20,
		"Roll": 101,
	}
	m3 := map[string]int{
		"Age":  21,
		"Roll": 100,
	}
	m4 := map[string]int{
		"Age":  20,
		"Roll": 100,
	}
	fmt.Println("Comparing m1 and m2 =>", maps.Equal(m1, m2))
	fmt.Println("Comparing m1 and m3 =>", maps.Equal(m1, m3))
	fmt.Println("Comparing m1 and m4 =>", maps.Equal(m1, m4))

	mySlice := []string{"A", "C", "D"}
	fmt.Println(mySlice)

	mySlice = slices.Insert(mySlice, 1, "B")
	fmt.Println(mySlice)

	mySlice = slices.Insert(mySlice, len(mySlice), "E", "F")
	fmt.Println(mySlice)

	deleteMethod()
}

func deleteMethod() {
	m := map[string]int{
		"A": 100,
		"B": 200,
		"C": 300,
		"D": 400,
		"E": 500,
	}

	fmt.Println("Original map =>", m)

	maps.DeleteFunc(m, func(k string, v int) bool {
		return v > 250
	})
	fmt.Println("Deleted keys for which values are more than 250 =>", m)
}
