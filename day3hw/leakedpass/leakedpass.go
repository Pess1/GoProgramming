package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
)

func checkPass(pass string, listname string) string {
	passlist, err := os.Open(listname)
	defer passlist.Close()

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(3)
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
	var listname string

	flag.StringVar(&listname, "l", "", "--l *password list file name*")
	flag.StringVar(&pass, "p", "", "--p *password to check if compromised*")
	flag.Parse()

	if pass == "" {
		fmt.Println("Please provide password")
		os.Exit(1)
	}

	if listname == "" {
		fmt.Println("Please provide the passwordlist file")
		os.Exit(2)
	}

	output := checkPass(pass, listname)

	fmt.Println(output)
}
