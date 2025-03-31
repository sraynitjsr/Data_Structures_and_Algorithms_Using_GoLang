// Final Commit Done

package hashset

import "fmt"

type MyHashSet struct {
	myMap map[interface{}]string
}

func (s *MyHashSet) add(val int) {
	_, presentOrNot := s.myMap[val]
	if !presentOrNot {
		s.myMap[val] = "Added"
	}
}

func (s *MyHashSet) remove(val int) {
	delete(s.myMap, val)
}

func (s *MyHashSet) display() {
	for value := range s.myMap {
		fmt.Print(value, " ")
	}
}

func Start_Optimized() {
	myMap := make(map[interface{}]string)
	myHashSet := &MyHashSet{
		myMap: myMap,
	}
	myHashSet.add(10)
	myHashSet.add(30)
	myHashSet.add(20)
	myHashSet.add(50)
	myHashSet.add(40)

	fmt.Print("First Run  => ")
	myHashSet.display()

	myHashSet.add(60)
	myHashSet.add(30)
	myHashSet.add(20)

	fmt.Print("\nSecond Run => ")
	myHashSet.display()

	myHashSet.remove(60)

	fmt.Print("\nThird Run => ")
	myHashSet.display()

	fmt.Println("")
}
