package main

import (
	"fmt"
)

// creating a typed struct
type thing struct {
	name  string
	value int
}

// creating a func that uses the typed struct "thing"
func update(things []thing) []thing {
	for i := range things {
		things[i].value *= 2
	}

	things = append(things, thing{name: "books", value: 100})
	return things //we have to return a slice if it is mutated in a function to update the
}
func main() {
	//creating a slice literal with composite literals for the struct element
	items := []thing{
		thing{name: "pens", value: 10},
		thing{name: "notepads", value: 10},
	}

	fmt.Println("Before:", items)
	items = update(items)
	fmt.Println("After:", items)
}
