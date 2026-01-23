package main

import (
	"fmt"
)

// 3. now below creating named types
type album3 struct {
	title string
}

type album4 struct {
	title string
}

func main() {
	var album = struct { //1. created an album struct type
		title  string
		artist string
		year   int
		copies int
	}{
		"Sifar",
		"Lucky Ali",
		1997,
		10000000,
	}
	fmt.Println(album)
	//above is not conveninet because
	//2. suppose i create two different struct types of anonymopus types but with the same structural similarity

	var album1 = struct {
		title string
	}{
		"Tanha Dil",
	}

	var album2 = struct {
		title string
	}{
		"Venga Boys",
	}
	album1 = album2
	fmt.Println(album1, album2) //this shows that these kind of structs can be assigned to one another irrespective of their anonamity

	//using named type structs from 3.

	var a3 = album3{
		"Bhoomi",
	}

	var a4 = album4{
		"Cactus",
	}
	//a3 = a4 //throws error as they are not the same named types but can be type converted as shown below
	a3 = album3(a4)
	fmt.Println(a3, a4)
}
