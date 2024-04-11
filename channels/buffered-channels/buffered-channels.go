package main

import "fmt"

func main() {
	channel := make(chan string, 2) // Buffer: 2
	channel <- "Hello World!"
	channel <- "Programming in Go!"
	/*
		The next line generates Deadlock by overflowing buffer (stackoverflow)
	*/
	// channel <- "Test deadlock!"

	msg := <-channel
	msg2 := <-channel
	/*
		The next line generates Deadlock by overflowing buffer (stackoverflow)
	*/
	// msg3 := <-channel
	fmt.Println(msg)
	fmt.Println(msg2)
	/*
		In case you want to test the DeadLock, Go demands that every variable needs to be used.
	*/
	// fmt.Println(msg3)
}
