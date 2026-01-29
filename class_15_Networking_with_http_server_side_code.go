// A basic program to create a basic web server which fetches user defined url path and displays it in the web page
package main

import (
	"fmt"
	"log"
	"net/http" //package used to deal with web server creation
)

// a func handler function that takes 2 declarations 1.reposnsewriter 2.http.request that takes requests from client and provide response
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! from %s\n", r.URL.Path[1:])
}

// In the main function
func main() {
	http.HandleFunc("/", handler)                //HandleFunc registers handler for all paths starting with /
	log.Fatal(http.ListenAndServe(":8080", nil)) //logs and exits if server fails to start
}
