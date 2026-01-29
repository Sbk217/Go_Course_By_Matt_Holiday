// This code acts as a client that retrieves a todo item from the JSONPlaceholder API and prints it locally
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// JSONPlaceholder is a free online REST API that you can use whenever you need some fake data
const url = "https://jsonplaceholder.typicode.com"

type todo struct { //the above typicode JSONPlaceholder provide 4 fields but you have the freedom to use just 3 in your JSON and ignore the ones you do not want
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Printf("%#v\n", item)
	}

}
