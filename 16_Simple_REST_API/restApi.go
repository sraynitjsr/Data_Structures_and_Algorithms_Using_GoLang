package simplerestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func paramsModify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["param1"]
	age := params["param2"]

	ageInInt, _ := strconv.ParseInt(age, 10, 64)

	emp := Employee{
		Age:  ageInInt,
		Name: name,
	}
	out, _ := json.Marshal(&emp)
	_, err := w.Write([]byte(out))
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

func Start() {
	fmt.Println("Starting Simple REST API.........")

	r := mux.NewRouter()
	r.HandleFunc("/{param1}/{param2}", paramsModify)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
