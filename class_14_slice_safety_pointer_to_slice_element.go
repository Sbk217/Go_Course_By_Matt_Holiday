package main

import (
	"fmt"
)

type user struct {
	name  string
	count int
}

func addTo(u *user) {
	u.count++
}

func main() {
	users := []user{{"alice", 0}, {"bob", 0}} //users records are created which is a slice having 2 users ':=' is used
	// to initialize a new variable and also infer the type of the variable in this instance a slice

	alice := &users[0]         //a new variable 'alice' is created and points to the 0th position of the slice
	amy := user{"amy", 0}      //a new variable 'amy' of type user struct is created here
	users = append(users, amy) //the above 'amy' user is appended to the pre-existing users variable
	//note here that alice still points to the old stale array where as users have been allocated a new array after the appending of amy
	//this is a risk of alloting pointers to an slice element
	addTo(alice)
	fmt.Println(users)
	fmt.Println("count what alice has: ", alice.count)

}
