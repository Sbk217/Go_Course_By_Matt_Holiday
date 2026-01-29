// a client server code that serve the local request by making a client request to the typicode REST API and then reformatting that response and providing it to us
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com"

var form = `
<h1>Todo #{{.ID}}</h1>
<div
`

type todo struct { //the above typicode JSONPlaceholder provide 4 fields but you have the freedom to use just 3 in your JSON and ignore the ones you do not want
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handler(w http.ResponseWriter, r *http.Request) {

}
