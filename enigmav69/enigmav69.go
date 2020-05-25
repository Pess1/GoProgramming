package main

import (
	"fmt"
	"flag"
	"strings"
)

func main() {
	var input string

	flag.StringVar(&input, "input", "a", "--input *String*, default value a lowercase only")

	flag.Parse()

	encrypt(input)	
}

func encrypt(input string) {
	letters := []rune(input)
	var alphabet = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var output = ""

	for _, letter := range letters {
		for i := 0; i < len(alphabet); i++ {
			if string(letter) == alphabet[i] {
				
				for x := 0; x <= i; x++ {
					output = output + "69"
				}
			}

		}
		output = output + "-"
	}

	decrypt(output)

	fmt.Println(output)
}

func decrypt (estring string) {
	letters := strings.Split(estring, "-")
	var alphabet = []string{"empty","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var dletters = ""

	for _, letter := range letters {
		var i int
		i = (len(strings.TrimSuffix(letter, "\n")) / 2)
		
		if alphabet[i] != "empty" {
			dletters = dletters + alphabet[i]
		} else {
			dletters = dletters + " "
		}
	}

	fmt.Println(dletters)
}
