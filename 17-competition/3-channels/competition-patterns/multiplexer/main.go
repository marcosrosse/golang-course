package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("Hello, world"), write("Programming  in GO!")) // write messages to two channels

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel) // receive a message from the channel
	}

}

func multiplexer(channelIn1, channelIn2 <-chan string) <-chan string { // receive from two channels, send to one channel
	channelOut := make(chan string)

	go func() {
		for {
			select {
			case message := <-channelIn1:
				channelOut <- message
			case message := <-channelIn2:
				channelOut <- message
			}
		}
	}()
	return channelOut
}

func write(text string) <-chan string {
	channel := make(chan string) // create a channel

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)            // send a message to the channel
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000))) // random sleep
		}
	}()
	return channel
}
