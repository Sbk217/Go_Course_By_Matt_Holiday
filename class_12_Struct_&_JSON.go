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
		2,
		nil,
		time.Now(),
	}
	fmt.Printf("%T %+[1]v\n", e2)

	//4 calling e3 after initializing giving explicit field names in the
	//initializer
	var e3 = Employee{
		Name:   "Rob",
		Number: 3,
		boss:   e2.boss,
		Hired:  time.Now(),
	}
	fmt.Printf("%T %+[1]v\n", e3)

	//creating a company c a map of strings to employee pointers
	c := map[string]*Employee{}
	//here Tobias is the boss of the company c
	//here a very important concept of maps, pointers and structs is showcased
	//here if we used c:= map[string]Employee then we cannot set the following such as
	//c["Tobias"] = &Employee as the map will point to a struct type map memory allocation is
	//dynamic which assignes memory as per the data given in the map which will result in a
	//dangling pointer for c["Tobias"] thus we provide the address for the struct employee
	//ie a map having a struct pointer not a simple struct

	c["Tobias"] = &Employee{
		Name:   "Tobias",
		Number: 4,
		boss:   e2.boss,
		Hired:  time.Now(),
	}
	fmt.Printf("%T %+[1]v\n", c["Tobias"])

	c["Stefan"] = &Employee{
		Name:   "Stefan",
		Number: 1,
		boss:   c["Tobias"],
		Hired:  time.Now(),
	}
	//will show the boss but since it is a pointer only the address is shown
	//in the output
	fmt.Printf("%T %+[1]v\n", c["Stefan"])

}
