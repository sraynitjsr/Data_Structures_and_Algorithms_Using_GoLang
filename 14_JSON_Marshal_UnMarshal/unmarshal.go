package jsonmarshalunmarshal

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Price string `json:"price"`
	Year  string `json:"year"`
}

type Product struct {
	Name string `json:"name"`
	Info Info   `json:"info"`
}

func unmarshal() {
	jsonData := `{
		"name": "Subhradeep",
		"info": {
			"price": "100",
			"year": "2020"
		}
	}`

	var structData Product
	err := json.Unmarshal([]byte(jsonData), &structData)
	if err != nil {
		fmt.Println(structData)
	}
}

/*
	// Below UnMarshalling is when we don't know JSON content

	jsonData := `{
		"name": "Subhradeep",
		"info": {
			"price": "1111",
			"year": "2222"
		}
	}`
	var structData interface{}
	json.Unmarshal([]byte(jsonData), &structData)
	fmt.Println(reflect.TypeOf(structData), structData)
*/
