package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1] //this is 2 index slice operator

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := b[0:2] //WTF moment as how can this be possible

	fmt.Println("c = ", c) //how can c show [1 2] ??
	// because When you slice an array/slice, the new slice still "sees" the entire underlying array.

	d := a[0:1:1] // [i:j:K] len= j-i and cap=k-i
	fmt.Println("d = ", d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	e := d[0:2]
	fmt.Println("e = ", e) //fails as d has only capacity = 1

	// when using 2 index slice operator the capacity of the slice is the
	// same as the capacity of the underlying array
	// when using 3 index slice operator the capacity of the slice can be
	// controlled just as the length of the 3 indexed slice
}
