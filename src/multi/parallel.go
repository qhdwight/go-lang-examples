package main

import (
	"fmt"
	"time"
)

func say() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello world!")
}

func main() {
	go say()
	fmt.Println("Goodbye world! Or?")
	time.Sleep(200 * time.Millisecond)
}
