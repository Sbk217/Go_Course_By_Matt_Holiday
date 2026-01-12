package main

import (
	"fmt"
)

func main_test() {
	s := "elite"

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
}
