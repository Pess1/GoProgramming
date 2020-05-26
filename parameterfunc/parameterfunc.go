package main

import (
	"fmt"
	"flag"
)

func calculateTotal(squat int, bench int, dl int) int {
	total := squat+bench+dl
	return total
}

func main() {
	var squat int
	var bench int
	var dl int

	flag.IntVar(&squat, "s", 0, "--s = squat")
	flag.IntVar(&bench, "b", 0, "--b = bench")
	flag.IntVar(&dl, "d", 0, "--d = deadlift")

	flag.Parse()

	var output int

	output = calculateTotal(squat, bench, dl)

	fmt.Println(output)
}
