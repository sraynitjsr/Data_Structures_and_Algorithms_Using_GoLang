package newstuff

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func Start() {

	myString := "a good   example"
	fmt.Println("Trimmed String =>", strings.Fields(myString))

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

	// Another usage of maps.DeleteFunc()
	mm := map[string]int{
		"a": 2020,
		"f": 8080,
		"e": 3030,
		"d": 6060,
		"c": 4040,
	}

	fmt.Println("Original map =>", mm)

	fmt.Println("ASCII values :-", "a =>", int('a'), " || f =>", int('f'), " || e =>", int('e'), " || d =>", int('d'), " || c =>", int('c'))

	maps.DeleteFunc(mm, func(k string, v int) bool {
		return int(k[0]) < 100
	})
	fmt.Println("Deleted keys for which ASCII values are less than 100 =>", mm)
}
