package jsonmarshalunmarshal

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func marshal() {
	author := Author{
		Name: "Sachin",
		Age:  20,
	}
	book := Book{
		Title:  "DSA",
		Author: author,
	}
	byteArr, _ := json.MarshalIndent(book,"","  ")
	fmt.Println(string(byteArr))
}

