package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	_ "image/png"
	_ "image/jpeg"
	_ "image/gif"
)

func convert(file1 string, file2 string) {
	f1, err1 := os.Open(file1)
	f2, err2 := os.Open(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error file1:", err1, "Error file2:", err2)
		os.Exit(2)
	}

	defer f1.Close()
	defer f2.Close()

	scanner1 := bufio.NewScanner(f1)
	content1 := ""
	for scanner1.Scan() {
		content1 += scanner1.Text()
	}

	scanner2 := bufio.NewScanner(f2)
	content2 := ""
	for scanner2.Scan() {
		content2 += scanner2.Text()
	}

	maxlen := 0
	minlen := 0
	if len(content1) <= len(content2) {
		maxlen = len(content2)
		minlen = len(content1)
	} else {
		maxlen = len(content1)
		minlen = len(content2)
	}

	var counter float64
	counter = 0
	bytes1 := []byte(content1)
	bytes2 := []byte(content2)
	for i := 0; i < minlen; i++ {
		if bytes1[i] == bytes2[i] {
			counter += 100
		}
	}

	var output float64
	output = counter/float64(maxlen)

	fmt.Println("The files match:", output, "%")

}

func main() {
	var file1 string
	var file2 string

	flag.StringVar(&file1, "f1", "", "--f1 = *file name*")
	flag.StringVar(&file2, "f2", "", "--f2 = *file name*")
	flag.Parse()

	if file1 == "" || file2 == "" {
		fmt.Println("Please provide both files")
		os.Exit(1)
	}

	convert(file1, file2)
}
