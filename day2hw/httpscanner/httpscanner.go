package main

import (
	"fmt"
	"net/http"
	"flag"
	"os"
	"bufio"
)

func sendRequest(url string) (string, string) {
	var body string
	var status string

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("There was an unexpected error:", err)
		os.Exit(2)
	}

	defer resp.Body.Close()

	status = "Response status: " + resp.Status + "\n"

	bodyscanner := bufio.NewScanner(resp.Body)
	for bodyscanner.Scan() {
		body = body + bodyscanner.Text() + "\n"
	}

	fmt.Println(resp.Header)

	return status, body
}

func main () {
	var url string
	var body string
	var status string

	flag.StringVar(&url, "u", "", "--u *url*")
	flag.Parse()

	if url == "" {
		fmt.Println("Please provide an url")
		os.Exit(1)
	} else {
		status, body = sendRequest(url)
	}

	fmt.Println(status, body)
}
