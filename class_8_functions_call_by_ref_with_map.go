package main

//code that uses the concept of function
//to assign an slice and do a function
//job and also return a value use concept of
//pass by reference
import (
	"fmt"
)

func do_map(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int) //m1 is a local variable in the function
	m1[4] = 4              //does not change m in any way
	fmt.Println("m1:", m1)
}
func main_test() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	do_map(m)
	fmt.Println("m:", m)
}
