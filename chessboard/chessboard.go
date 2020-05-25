package main

import "fmt"

func main() {
	
	for i := 0; i < 9; i++ {
		var word string
		for x := 0; x < 9; x++ {
			if (x%2) == 0 {
				word = word + " # "
			} else {
				word = word + " + "
			}
		}
		fmt.Println(word)
	}
}
