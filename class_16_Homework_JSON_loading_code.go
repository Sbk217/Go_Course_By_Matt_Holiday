// question link from Matt Holidays course : https://github.com/matt4biz/go-class-slides/blob/trunk/xmas-2020/go-16-hw3-slides.pdf
// Brief read through the JSON interface of a popular cartoon series and download each url once and build an offline index write a tool xkcd, that uses the index
// prints the URL, date and title of each comic whose title or title matches a list of search terms provided on the command line.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// getOne returns the metadata for one comic by number
func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	//check if the page with the url is even reachable by internet or not
	if err != nil {
		fmt.Fprintf(os.Stderr, "Stopped reading: %s\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	//checking if the page with the url is providing a vaild 200 http status code or not
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Skipping %d: got %d\n", i, resp.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "bad Body: %s\n", err)
		os.Exit(-1)
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 { //check if a command line argument is provided or not by checking the count
		output, err = os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	//the below output will be in JSON format so adding the front and end "[" and "]"

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	//stop if we get 2 404's
	for i := 1; fails < 2; i++ {

		if data = getOne(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",") //OB1
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}

		fails = 0
		cnt++

	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
