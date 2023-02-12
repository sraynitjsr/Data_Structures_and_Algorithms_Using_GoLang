package hashset

import "fmt"

type set struct {
	data []int
}

func GetHashSet() set {
	emptySet := set{
		data: []int{},
	}
	return emptySet
}

func (temp *set) add(val int) {
	for index, element := range temp.data {
		if element == val {
			fmt.Printf("\nRepetitive elements not allowed in Set, found %d at index %d\n", element, index)
			return
		}
	}
	temp.data = append(temp.data, val)
}

func (temp *set) remove(val int) {
	newSet := []int{}
	for _, element := range temp.data {
		if element != val {
			newSet = append(newSet, element)
		}
	}
	temp.data = newSet
}

func Start() {
	fmt.Println("Inside HashSet")

	myHashSet := set{
		data: []int{},
	}

	myHashSet.add(10)
	myHashSet.add(-20)
	myHashSet.add(30)
	myHashSet.add(40)
	myHashSet.add(-50)

	fmt.Println("\nCurrent Set is", myHashSet.data)

	myHashSet.add(30)
	myHashSet.add(60)
	myHashSet.add(70)
	myHashSet.add(80)

	myHashSet.remove(-50)

	fmt.Println("\nCurrent Set is", myHashSet.data)
}

