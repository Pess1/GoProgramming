package main

import "fmt"

func main() {

	var i int
	
	var total int

	for i <= 1000 {
		total = total + i
		i++
	}
	
	fmt.Println(total)
}
