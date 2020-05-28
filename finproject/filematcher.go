package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"image"
	"bytes"
	"image/png"
	_ "image/jpeg"
	_ "image/gif"

	"github.com/nfnt/resize"
	//^Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>
)

//Resizes the images much smaller and converts the images to strings
func convertImg(file1 string, file2 string) {
	imgFile1, err1 := os.Open(file1)
	imgFile2, err2 := os.Open(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error opening files:", err1, err2)
		os.Exit(2)
	}

	defer imgFile1.Close()
	defer imgFile2.Close()

	//Decode the opened files to images
	f1, _, dErr1 := image.Decode(imgFile1)
	f2, _, dErr2 := image.Decode(imgFile2)

	if dErr1 != nil || dErr2 != nil {
		fmt.Println("Decoding failed:", dErr1, dErr2)
		os.Exit(2)
	}

	//Resize the images to be smaller
	f1 = resize.Resize(50, 25, f1, resize.MitchellNetravali)
	f2 = resize.Resize(50, 25, f2, resize.MitchellNetravali)

	//Encode the images back to bytes
	buf1 := new(bytes.Buffer)
	encodeErr1 := png.Encode(buf1, f1)
	img1Bytes := buf1.Bytes()

	buf2 := new(bytes.Buffer)
	encodeErr2 := png.Encode(buf2, f2)
	img2Bytes := buf2.Bytes()

	if encodeErr1 != nil || encodeErr2 != nil {
		fmt.Println("Failed to convert image:", encodeErr1, encodeErr2)
		os.Exit(2)
	}

	//Convert the byte slices to strings since the compare function requires string
	content1 := ""
	for _,  byte1 := range img1Bytes {
		content1 += string(byte1)
	}

	content2 := ""
	for _, byte2 := range img2Bytes {
		content2 += string(byte2)
	}

	//Call compare function
	compare(content1, content2)
}

//Converts files to strings
func convert(file1 string, file2 string) {
	f1, err1 := os.Open(file1)
	f2, err2 := os.Open(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error file1 opening files:", err1, err2)
		os.Exit(2)
	}

	defer f1.Close()
	defer f2.Close()

	//Scan the file content and turn them into strings
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

	//Call compare function
	compare(content1, content2)

}

//Compares both strings to each other 1 symbol at a time
func compare(content1 string, content2 string) {
	//Establish a counter that counts the times same symbols appear at same index
	var counter float64
	counter = 0

	//Turn the string to byte slices to iterate through the content
	bytes1 := []byte(content1)
	bytes2 := []byte(content2)
	length1 := len(bytes1)
	length2 := len(bytes2)
	fmt.Println("Symbol amount file1:", length1 ,"Symbol amount file2:", length2)

	for i := 0; i < length1 && i < length2; i++ {
		if bytes1[i] == bytes2[i] {
			counter += 100
		}
	}

	//Get the divider to calculate percentage
	maxlen := 0
	if length1 <= length2 {
		maxlen = length2
	} else {
		maxlen = length1
	}

	//By using float64 we can get a really accurate percentage
	var output float64
	output = counter/float64(maxlen)

	fmt.Println("The files match:", output, "%")

}

func main() {
	var file1 string
	var file2 string
	var isImg bool

	flag.StringVar(&file1, "f1", "", "--f1 = *file name*")
	flag.StringVar(&file2, "f2", "", "--f2 = *file name*")
	flag.BoolVar(&isImg, "i", false, "--i, enable if you want to compare images")
	flag.Parse()

	fmt.Println("=============================\n===FFFFFFF=MMMM=======MMMM===\n===FF======MM=MM=====MM=MM===\n===FFFFFFF=MM==MM===MM==MM===\n===FF======MM===MM=MM===MM===\n===FF======MM=====M=====MM===\n===FF======MM===========MM===\n=============================")
	//If file is missing program exits, if --i is written img mode is enabled, otherwise program uses normal mode
	if file1 == "" || file2 == "" {
		fmt.Println("Please provide both files")
		os.Exit(1)
	} else if isImg {
		fmt.Println("Image mode engaged!!!!")
		convertImg(file1, file2)
	} else {
		fmt.Println("Text mode engaged!!!!")
		convert(file1, file2)
	}
}
