// Dont mistake the concept here used with ovveride we dont do ovveride here in Go which is of inheritance and here PairWithLength is not a subclass of Pair
package main

import (
	"fmt"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is: %s", p.Path, p.Hash)
}

func (pw PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is: %s; and length is: %d", pw.Path, pw.Hash, pw.Length)
}

type PairWithLength struct {
	Pair
	Length int
}

func main() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xded"}, 133}

	fmt.Println(p)
	fmt.Println(pl)
}
