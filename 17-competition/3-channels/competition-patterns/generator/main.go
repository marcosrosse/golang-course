package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("My channel generator") // write a message to the channel write

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel) // receive a message from the channel
	}

}

func write(text string) <-chan string {
	channel := make(chan string) // create a channel

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text) // send a message to the channel
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
