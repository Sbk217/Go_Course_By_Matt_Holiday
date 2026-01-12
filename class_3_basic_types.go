package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func test() {
	a := 2
	b := 2.1
	c := 3.1
	d := 4

	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)
	fmt.Printf("c: %T %[1]v\n", c)
	d = int(c)
	fmt.Printf("d: %T %[1]v\n", d)
	fmt.Printf("c: %T %[1]v\n", c)

}
