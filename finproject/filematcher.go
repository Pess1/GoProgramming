package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"image"
	"bytes"
	_ "image/png"
	 "image/jpeg"
	_ "image/gif"

	"github.com/nfnt/resize"
	//Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>
)

//Converts images to string
func convertImg(file1 string, file2 string) {
	imgFile1, err1 := os.Open(file1)
	imgFile2, err2 := os.Open(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error opening files:", err1, err2)
		os.Exit(2)
	}

	defer imgFile1.Close()
	defer imgFile2.Close()

	f1, _, dErr1 := image.Decode(imgFile1)
	f2, _, dErr2 := image.Decode(imgFile2)

	if dErr1 != nil || dErr2 != nil {
		fmt.Println("Decoding failed:", dErr1, dErr2)
		os.Exit(2)
	}

	f1 = resize.Resize(50, 0, f1, resize.MitchellNetravali)
	f2 = resize.Resize(50, 0, f2, resize.MitchellNetravali)

	buf1 := new(bytes.Buffer)
	encodeErr1 := jpeg.Encode(buf1, f1, nil)
	img1Bytes := buf1.Bytes()

	buf2 := new(bytes.Buffer)
	encodeErr2 := jpeg.Encode(buf2, f2, nil)
	img2Bytes := buf2.Bytes()

	if encodeErr1 != nil || encodeErr2 != nil {
		fmt.Println("Failed to convert image:", encodeErr1, encodeErr2)
		os.Exit(2)
	}

	content1 := ""
	for _,  byte1 := range img1Bytes {
		content1 += string(byte1)
	}

	content2 := ""
	for _, byte2 := range img2Bytes {
		content1 += string(byte2)
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

	compare(content1, content2, maxlen, minlen)
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

	compare(content1, content2, minlen, maxlen)

}

//Compares both strings to each other 1 symbol at a time
func compare(content1 string, content2 string, minlen int, maxlen int) {
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
	var isImg bool

	flag.StringVar(&file1, "f1", "", "--f1 = *file name*")
	flag.StringVar(&file2, "f2", "", "--f2 = *file name*")
	flag.BoolVar(&isImg, "i", false, "--i, enable if you want to compare images")
	flag.Parse()

	if file1 == "" || file2 == "" {
		fmt.Println("Please provide both files")
		os.Exit(1)
	} else if isImg {
		fmt.Println("Image mode engaged!!!!")
		convertImg(file1, file2)
	}

	convert(file1, file2)
}
