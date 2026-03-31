// Dont mistake the concept here used with ovveride we dont do ovveride here in Go which is of inheritance and here PairWithLength is not a subclass of Pair
package main

import (
	"fmt"
	"path/filepath"
)

type Pair2 struct {
	Path string
	Hash string
}

func (p Pair2) String() string {
	return fmt.Sprintf("Hash of %s is: %s", p.Path, p.Hash)
}

func (pw PairWithLength2) String() string {
	return fmt.Sprintf("Hash of %s is: %s; and length is: %d", pw.Path, pw.Hash, pw.Length)
}

type PairWithLength2 struct {
	Pair2
	Length int
}

func Filename(p Pair2) string {
	return filepath.Base(p.Path)
}

func main() {
	p := Pair2{"/usr", "0xfdfe"}
	pl := PairWithLength2{Pair2{"/usr/lib", "0xded"}, 133}

	fmt.Println(p)
	fmt.Println(pl)

	fmt.Println(Filename(p)) //This WORKS !!!!

	// fmt.Println(Filename(pl)) //This DOES NOT WORK !!!! because there is not concept of ovveride or inheritance in Go Land

	fmt.Println(Filename(pl.Pair2)) //This WORKS AGAIN!!!! because we are using the concept of composition/promoting here
}
