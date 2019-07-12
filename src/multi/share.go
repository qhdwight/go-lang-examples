package main

import "fmt"

func sliceSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Send sum to channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	l := len(s)

	c := make(chan int, 1)
	go sliceSum(s[:l/2 ], c)
	go sliceSum(s[ l/2:], c)
	x, y := <-c, <-c // Wait until there is a value from channel

	fmt.Printf("%d, %d, %d\n", x, y, x+y)
}
