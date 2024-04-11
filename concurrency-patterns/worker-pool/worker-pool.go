package main

import "fmt"

// The idea is that the next available worker will be processing the next number of fibonnaci
// worker-1 do fibonacci(0) | worker-1 is now busy
// worker-2 do fibonacci(1) | worker-2 is now busy
// worker-3 do fibonacci(2) | worker-3 is now busy
// ... and so on until 45, the workers will be taking turns when available

func main() {
	tasksChannel := make(chan int, 45)
	resultsChannel := make(chan int, 45)

	createWorkerPool(tasksChannel, resultsChannel, 5)

	for i := 0; i < 45; i++ {
		tasksChannel <- i // Sends data to the channel and the next available worker will perform the operation
	}

	close(tasksChannel)

	for i := 0; i < 45; i++ {
		result := <-resultsChannel // Prints the result as soon as it receives data on the channel
		fmt.Println(result)
	}
}

/*
tasksChannel is a read-only channel (<-chan)
resultsChannel is a write-only channel (chan<-)
workers are the number of concurrent functions running
*/
func createWorkerPool(tasksChannel <-chan int, resultsChannel chan<- int, workers int) {
	for i := 0; i <= workers; i++ {
		go worker(tasksChannel, resultsChannel)
	}
}

func worker(tasksChannel <-chan int, resultsChannel chan<- int) {
	for number := range tasksChannel {
		resultsChannel <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
