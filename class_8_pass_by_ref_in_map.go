package main

//code that uses the concept of function
//to assign an map and do a function call
//job which alters the original value
//via ref (pointers) showcasing concept of
//pass by reference

import (
	"fmt"
)

func do_map_ref(m1 *map[int]int) {
	(*m1)[3] = 0
	*m1 = make(map[int]int) //m1 is a local variable in the function
	(*m1)[4] = 4            //does not change m in any way
	fmt.Println("m1:", m1)
}
func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	do_map_ref(&m) //need to pass the address of the map
	fmt.Println("m:", m)
}
