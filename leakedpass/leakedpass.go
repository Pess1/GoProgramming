package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
)

func checkPass(pass string) string {
	passlist, err := os.Open("10-million-password-list-top-1000000.txt")
	defer passlist.Close()

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(2)
	}

	output := "Password not found on the password list"
	scanner := bufio.NewScanner(passlist)
	for scanner.Scan() {
		if pass == scanner.Text() {
			output = pass + " has been found from the password list"
			return output
		} 
	}

	return output
}

func main() {
	var pass string
	output := ""

	flag.StringVar(&pass, "p", "", "--p *password to check if compromised*")
	flag.Parse()

	if pass == "" {
		fmt.Println("Please provide password")
		os.Exit(1)
	}

	output = checkPass(pass)

	fmt.Println(output)
}
