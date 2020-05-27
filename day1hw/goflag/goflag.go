package main

import (
	"fmt"
	"flag"
)

func main() {
	var amount int
	var name string

	flag.IntVar(&amount, "amount", 1, "--amount **Number**, default 1")
	flag.StringVar(&name, "name", "empty", "--name **String**")

	flag.Parse()

	if name == "empty" {
		fmt.Println("Who is you! Identify yourself!")
	} else {

		for i:=0; i < amount; i++ {
			fmt.Println("Whassap", name)
		}
	}
}
