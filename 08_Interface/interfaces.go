package interfaces

import (
	"fmt"

	"github.com/sraynitjsr/08_Interface/embedding_example"
)

func Start() {
	fmt.Println("Inside Interfaces")

	fmt.Println("\nSimple Interface")
	SimpleInterface()

	fmt.Println("\nEmbedding and Loose Coupling")
	embedding_example.Embedding()
}
