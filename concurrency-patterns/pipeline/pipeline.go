package main

import "fmt"

/*
Pipeline:
[start] >>>>>>>>> [Stage 1] >>>>>>>>> [Stage 2] >>>>>>>>> [Stage 3] >>>>>>>>> [end]

Since all channels are unbuffered, all interactions happen synchronously!
- Stage 1 creates a dataChannel that interates throughout inputs sending 1 by 1 to the channel.
- Stage 2 receives the dataChannel from Stage 1, and wait it to have a value to perform the action.
- Stage 3 receives the finalChannel from Stage 2, and also wait it to have a value to print it.
*/
func main() {
	// input
	numbers := []int{2, 3, 4, 7, 1}
	// Stage 1
	dataChannel := sliceToChannel(numbers)
	// Stage 2
	finalChannel := sq(dataChannel)
	// Stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
