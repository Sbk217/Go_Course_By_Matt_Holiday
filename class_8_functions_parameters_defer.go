package main

//code that uses the concept of function
//to assign an array and do a function
//job and also return a value use concept of
//pass by value
import (
	"fmt"
)

func do(b [3]int) int {
	b[0] = 0
	return b[1]
}
func main_test() {
	a := [3]int{1, 2, 3}
	v := do(a)

	fmt.Println(a, v)
}
