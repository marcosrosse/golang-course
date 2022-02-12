package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500) // it runs in half a second
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2) // it run in 2 seconds
			channel2 <- "Channel 2"
		}
	}()

	for {
		select {
		case msgChannel1 := <-channel1: // if receive the channel 1 message, print it
			fmt.Println(msgChannel1)
		case msgChannel2 := <-channel2: // if receive the channel 2 message, print it
			fmt.Println(msgChannel2)

		}
	}

}
