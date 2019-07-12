package main

import (
	"fmt"
)

func fibonacci(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // Wait until we can send x to channel c
			x, y = y, x+y
		case <-quit: // Wait until there is a value from the quit channel
			fmt.Println("I have quit!")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go func() { // Anonymous lambda and execution
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibonacci(c, quit)
}
