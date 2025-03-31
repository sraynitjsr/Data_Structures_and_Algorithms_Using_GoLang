// Final Commit Done

package func_prog

import "fmt"

type MapperFunc func(int) int
type PredicateFunc func(int) bool
type ReducerFunc func(int, int) int

func Map(slice []int, mapper MapperFunc) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

func Filter(slice []int, predicate PredicateFunc) []int {
	var result []int
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce(slice []int, reducer ReducerFunc, initial int) int {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func Start() {
	fmt.Println("Functional Programming in GoLang")
	numbers := []int{1, 2, 3, 4, 5}

	// Example of mapping: double each element
	doubledNumbers := Map(numbers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubledNumbers)

	// Example of filtering: keep only even numbers
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)

	// Example of reducing: calculate the sum of all numbers
	sum := Reduce(numbers, func(acc, n int) int {
		return acc + n
	}, 0)
	fmt.Println("Sum of numbers:", sum)
}
