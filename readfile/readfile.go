package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

func read(input string) string {
	text, err := ioutil.ReadFile(input)

	if err != nil {
		fmt.Println("Error message:", err)
		return "Nothing to show"
	}

	return string(text)
}

func main() {
	var input string

	flag.StringVar(&input, "f", "", "--f *filename*")

	flag.Parse()

	output := read(input)

	fmt.Println(output)
}
