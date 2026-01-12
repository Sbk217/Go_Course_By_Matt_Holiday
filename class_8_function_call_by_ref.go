package main

//code that uses the concept of function
//to assign a slice and do a function
//job and also return a value use concept of
//pass by reference
import (
	"fmt"
)

func do_slice(b []int) int {
	b[0] = 0
	return b[1]
}
func main() {
	a := []int{1, 2, 3}
	v := do_slice(a)

	fmt.Println(a, v)
}
