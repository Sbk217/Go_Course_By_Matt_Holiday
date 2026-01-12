package main

import (
	"fmt"
)

func main_test() {
	a := []byte("string")
	fmt.Println(len(a), a)
	fmt.Println(a[2])
	fmt.Println(a[2:])
	fmt.Println(a[3:5], len(a[3:5]))
}
