package main

import "fmt"

// recursive functions

func fibonnaci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonnaci(position-2) + fibonnaci(position-1)
}

func worker(tasks <-chan int, results chan<- int) { // receive from tasks, send to results
	for number := range tasks {
		results <- fibonnaci(number)
	}

}

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}

}
