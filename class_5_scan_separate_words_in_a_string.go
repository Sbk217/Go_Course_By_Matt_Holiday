package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main_test() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int) //make initializes memory to the map
	//Without make, you'd have a nil map that
	//would panic if you tried to write to it

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv //ss is slice of kv struct

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")
	}

}
