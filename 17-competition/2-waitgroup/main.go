package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroups

func write(text string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Go Routine 1!")
		waitGroup.Done()
	}()

	go func() {
		write("Go Routine 2!")
		waitGroup.Done()
	}()

	go func() {
		write("Go Routine 3!")
		waitGroup.Done()
	}()

	go func() {
		write("Go Routine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}
