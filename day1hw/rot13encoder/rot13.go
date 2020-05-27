package main

import (
	"fmt"
	"flag"
)

func main() {
	var input string

	flag.StringVar(&input, "input", "test", "--input *text*, default value = test")

	flag.Parse()

	fmt.Println("Original string: ", input)

	letters := []rune(input)

	var output = ""

	for _, letter := range letters {
		rotvalue := rot(letter)
		output = output + string(rotvalue)
	}

	fmt.Println("Rotted string: ", output)
}

func rot(input rune) rune {
	if input >= 'a' && input <= 'z' {
		return 'a' + (input - 'a' + 13) % 26
	} else if input >= 'A' && input <= 'Z' {
		return 'A' + (input - 'A' + 13) % 26
	} else {
		return input
	}

}
