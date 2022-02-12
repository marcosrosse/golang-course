package main

import (
	"fmt"
	"time"
)

// channels

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // waiting for an message be send
		time.Sleep(time.Second)
	}
	close(channel) //close the channel
}

func main() {
	channel := make(chan string)
	go write("Hello, World!", channel) // goroutine to send an message to the channel, receiveid by the function write

	for message := range channel {
		fmt.Println(message)
	}
	fmt.Println("The end!")

}
