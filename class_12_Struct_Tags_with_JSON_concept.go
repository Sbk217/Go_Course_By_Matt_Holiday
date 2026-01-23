package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r := Response{Page: 1, Words: []string{"up", "in", "out"}}
	fmt.Printf("Internal representation: %#v\n", r) //internal representation

	//showing just JSON format
	j, _ := json.Marshal(r)                                     //converts go value r into json bytes using encoding/json package
	fmt.Println("byte representation: ", j)                     //will print just bytes
	fmt.Println("string converted representation: ", string(j)) //will print string converted

	var r2 Response
	_ = json.Unmarshal(j, &r2) //passing the byte slice to a pointer r2
	fmt.Printf("%#v\n", r2)

	//** NOTE: these concept of struct tags are extensivley used during SQL related code using GO lang to avoid bobby table effect and
	//sql injection.
}
