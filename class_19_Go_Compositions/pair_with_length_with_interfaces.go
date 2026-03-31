// We get around using the concept of inheritance via using interfaces in Go lang
package main

import (
	"fmt"
	"path/filepath"
)

type Pair3 struct {
	Path string
	Hash string
}

func (p Pair3) String() string {
	return fmt.Sprintf("Hash of %s is: %s", p.Path, p.Hash)
}

func (pw PairWithLength3) String() string {
	return fmt.Sprintf("Hash of %s is: %s; and length is: %d", pw.Path, pw.Hash, pw.Length)
}

type PairWithLength3 struct {
	Pair3  // ← embedded struct (not inherited)
	Length int
}

func (p Pair3) Filename2() string {
	return filepath.Base(p.Path)
}

// Creating interface here which takes the method Filename2
type Filenamer interface {
	Filename2() string
}

func main() {
	p := Pair3{"/usr", "0xfdfe"}
	var fn Filenamer = PairWithLength3{Pair3{"/usr/lib", "0xded"}, 133}

	fmt.Println(p.Filename2())
	fmt.Println(fn.Filename2()) // calls Pair3's Filename2(), promoted through embedding

}
