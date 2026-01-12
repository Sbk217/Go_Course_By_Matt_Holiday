package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func main() {
	f := fib()
	for i := f(); i < 100; i = f() {
		fmt.Println(i)
	}
}
