package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	boss   *Employee //pointer to another employee (a type can declare itself until and unless it doesn't throw a pointer)
	Hired  time.Time
}

func main() {
	var e, e1 Employee
	//multiple ways to fetch the var e here
	//1 calling e without initializing anything will return in nil and deault 0 values
	fmt.Printf("%T %+[1]v\n", e)

	//2 calling e1 after initializing the values
	e1.Name = "Samule"
	e1.Number = 1
	e1.Hired = time.Now()
	fmt.Printf("%T %+[1]v\n", e1)

	//3 calling e2 after initializing without giving explicit field names in the
	//initializer
	var e2 = Employee{
		"Sam",
		1,
		nil,
		time.Now(),
	}
	fmt.Printf("%T %+[1]v\n", e2)

	//4 calling e3 after initializing giving explicit field names in the
	//initializer
	var e3 = Employee{
		Name:   "Rob",
		Number: 1,
		boss:   e2.boss,
		Hired:  time.Now(),
	}
	fmt.Printf("%T %+[1]v\n", e3)
}
