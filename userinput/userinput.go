package main

import "fmt"

func main() {

	fmt.Print("Enter name: ")
	var name string
	fmt.Scanln(&name)
	
	fmt.Print("Enter age: ")
	var age string
	fmt.Scanln(&age)

	fmt.Println(name + age)
}
