package main

import (
	"fmt"
)

func main() {
	s := "elite"

	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s))
}
