package main

import (
	"fmt"
	"flag"
	"strings"
)

func main() {
	var input string
	var b bool

	flag.StringVar(&input, "input", "a", "--input *String*, default value a lowercase only")
	flag.BoolVar(&b, "b", false, "--b, type this to decrypt")

	flag.Parse()

	if b {
		decrypt(input)
	} else {
		encrypt(input)
	}
}

func encrypt(input string) {
	letters := []rune(input)
	var output = ""

	for _, letter := range letters {
		if string(letter) == " " {
			output = output + "-"
		} else {
			for x := 0; x <= int(letter); x++ {
				output = output + "69"
			}
			output = output + "."
		}

	}

	fmt.Println(output)
}

func decrypt (estring string) {
	words := strings.Split(estring, "-")
	var output = ""
	var c = ""

	for _, word := range words {
		chars := strings.Split(word, ".")
		for _, char := range chars {
			c = strings.TrimSuffix(char, "\n")
			output = output + string(rune(len(c) / 2 - 1))
		}
		output = output + " "
	}

	fmt.Println(output)

}
