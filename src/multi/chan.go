package main

import (
	"fmt"
	"time"
)

func main() {
	msgChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		msgChan <- "ping" // Send to channel
	}()

	msg := <-msgChan // Wait for channel value
	fmt.Println(msg)
}
