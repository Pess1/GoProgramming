package main

import "fmt"

func main() {
	fmt.Print("Enter number of loops: ")
	var amount int
	fmt.Scanln(&amount)
	
	fmt.Print("Enter what you want to say: ")
	var word string
	fmt.Scanln(&word)

	for i := 0; i < amount; i++ {

		if (i == 1) || (i == 11) || (i == 12) || (i == 13) || (i == 14) {
			fmt.Println("The current value is a symbol of white supremacy", i)
		} else {
			fmt.Println(word)
		}

	}
}
