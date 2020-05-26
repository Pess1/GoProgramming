
package main

import (
	"fmt"
	"flag"
)

type Point struct {
	Xcoord int
	Ycoord int
}

func buildTable(p Point) string {
	output := ""

	for i:=1; i <= 4; i++ {
		if p.Ycoord != i {
			output = output + "# # # #"
		} else {
			for x := 1; x <= 4; x++ {
				if x != p.Xcoord {
					output = output + "# "
				} else {
					output = output + "+ "
				}
			}
		}

		output = output + "\n"
	}

	return output
}

func main() {
	var p Point
	var x int
	var y int

	flag.IntVar(&x, "x", 1, "--x *value* 1-4")
	flag.IntVar(&y, "y", 1, "--y *value* 1-4")

	flag.Parse()

	p.Xcoord = x
	p.Ycoord = 5-y

	output := buildTable(p)
	fmt.Println(output)
}
