package main

import (
	"fmt"
	"runtime"
)

func getos() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	youros := ""
	yourarch := ""

	if os == "windows" {
		youros = "Windows"
	} else if os == "darwin" {
		youros = "MAC OS"
	} else if os == "linux" {
		youros = "Linux"
	} else {
		youros = "Unknown operating system"
	}

	if arch == "amd64" {
		yourarch = "64-bit"
	} else if arch == "386" {
		yourarch = "32-bit"
	}

	output := yourarch + " " + youros

	return output
}

func main() {
	fmt.Println("This program is running on:", getos())
}
