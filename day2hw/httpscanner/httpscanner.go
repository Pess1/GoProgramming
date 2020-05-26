package main

import (
	"fmt"
	"net/http"
	"flag"
	"os"
	"bufio"
)

func sendRequest(url string) (string, string) {
	var output string
	var status string

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("There was an unexpected error:", err)
		os.Exit(2)
	}

	defer resp.Body.Close()

	status = "Response status: " + resp.Status

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		output = output + scanner.Text() + "\n"
	}

	return status, output
}

func main () {
	var url string
	var output string
	var status string

	flag.StringVar(&url, "u", "", "--u *url*")
	flag.Parse()

	if url == "" {
		fmt.Println("Please provide an url")
		os.Exit(1)
	} else {
		status, output = sendRequest(url)
	}

	fmt.Println(status, output)
}
