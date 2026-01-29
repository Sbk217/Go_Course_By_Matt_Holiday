// This is a proxy server that fetches TODO data from an external API and renders it as HTML.
package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type todo1 struct { //the above typicode JSONPlaceholder provide 4 fields but you have the freedom to use just 3 in your JSON and ignore the ones you do not want
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.Id}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler1(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	resp, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable) //http.Error()->method, w->reponse write, err.Error()->string from the error, status code
		return                                                    //I have to return because say there is nothing in the server end to fetch from
	}

	defer resp.Body.Close() //if i did not get an error i have to close the body otherwise my server is going to run out of sockets

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var item todo1

	err = json.Unmarshal(body, &item) ///unmarshalling the JSOn after which it will be inside the GO struct

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine") //creating template to utilize the saved GO struct data

	tmpl.Parse(form)
	tmpl.Execute(w, item) //Template engine substitutes struct fields into the HTML template
	//concept of reflection is showcase in this, reflection inspects struct fields at runtime
}

func main() {
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}
