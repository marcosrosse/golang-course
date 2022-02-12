package main

import (
	"fmt"
	"time"
)

// competition

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("Hello, world!")
	write("Show me something!")
}
