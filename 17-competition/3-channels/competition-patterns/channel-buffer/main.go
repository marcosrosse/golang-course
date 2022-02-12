package main

// channels buffer
// a channel is always used in a split functions

import "fmt"

func main() {
	channel := make(chan string, 1) // to use a channel in a single function, you need to inform a buffer

	channel <- "Hello, world!"

	message := <-channel
	fmt.Println(message)
}
