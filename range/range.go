package main

import "fmt"

func main() {

	var taivaankappaleet = []string{"Maa", "Merkurius", "Mars", "Pluto"}
	var planets []string = taivaankappaleet[0:3]

	for _, planet := range planets {
		fmt.Println(planet)
	}
}
