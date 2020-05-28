package main

import (
	"fmt"
	"time"
)

func hello(ch chan int) {
	fmt.Println("=== hello() has started ===")
	time.Sleep(1*time.Second)
	fmt.Println("=== hello() stopped ===")
	ch <- 1
}

func hi(ch chan int) {
	fmt.Println("=== hi() has started ===")
	time.Sleep(2*time.Second)
	fmt.Println("=== hi() stopped ===")
	ch <- 2
}

func sup(ch chan int) {
	fmt.Println("=== sup() has started ===")
	fmt.Println("=== sup() stopped ===")
	ch <- 3
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	fmt.Println("=== main() has started ===")

	go hello(ch1)
	go hi(ch2)
	go sup(ch3)

	<- ch1
	<- ch2
	<- ch3

	fmt.Println("=== main() stopped ===")
}
