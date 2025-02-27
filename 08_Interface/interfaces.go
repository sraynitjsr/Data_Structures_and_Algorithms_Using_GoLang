package interfaces

import "fmt"

func Start() {
	fmt.Println("Inside Interfaces")

	fmt.Println("\nSimple Interface")
	SimpleInterface()

	fmt.Println("\nEmbedding and Loose Coupling")
	Embedding()
}
