package sorting

import (
	"fmt"
	"sort"
)

func normalSort() {
	myIntSlice := []int{1, 3, 2}

	fmt.Println("Original", myIntSlice)
	sort.Slice(myIntSlice, func(i, j int) bool {
		return myIntSlice[i] < myIntSlice[j]
	})
	fmt.Println("Ascending", myIntSlice)
	sort.Slice(myIntSlice, func(i, j int) bool {
		return myIntSlice[i] > myIntSlice[j]
	})
	fmt.Println("Descending", myIntSlice)
}

type Student struct {
	Name string
	Roll int
}

func customSort() {
	student1 := Student{
		Name: "A",
		Roll: 10,
	}
	student2 := Student{
		Name: "C",
		Roll: 30,
	}
	student3 := Student{
		Name: "B",
		Roll: 20,
	}
	students := []Student{student1, student2, student3}
	sort.Slice(students, func(i, j int) bool {
		return students[i].Roll > students[j].Roll
	})
	fmt.Println("Original", students)
	sort.Slice(students, func(i, j int) bool {
		return students[i].Roll < students[j].Roll
	})
	fmt.Println("Ascending", students)
	sort.Slice(students, func(i, j int) bool {
		return students[i].Roll > students[j].Roll
	})
	fmt.Println("Descending", students)
}

func Start() {
	fmt.Println("Inside Sorting")
	normalSort()
	customSort()
}
