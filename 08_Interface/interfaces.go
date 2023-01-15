package interfaces

import "fmt"

type Men struct {
	hobby string
}

func (temp *Men) getHobby() string {
	return temp.hobby
}

type Women struct {
	hobby string
}

func (temp *Women) getHobby() string {
	return temp.hobby
}

type MyInterface interface {
	getHobby() string
}

func getHobby(myInterface MyInterface) string {
	return myInterface.getHobby()
}

func Start() {
	fmt.Println("Inside Interfaces")
	man := Men{
		hobby: "Dancing",
	}
	woman := Women{
		hobby: "Singing",
	}

	fmt.Println("Without Interface")
	fmt.Print("\tMan's Hobby ", man.getHobby())
	fmt.Println("\tWoman's Hobby", woman.getHobby())

	fmt.Println("With Interface")
	fmt.Print("\tMan's Hobby ", getHobby(&man))
	fmt.Println("\tWoman's Hobby", getHobby(&woman))
}
