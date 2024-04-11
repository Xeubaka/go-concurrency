package main

import (
	"fmt"
	"time"
)

func main() {
	doneChannel := make(chan bool)

	go writeOnTerminal(doneChannel)

	time.Sleep(time.Second * 3)

	// doneChannel <- true // This also work, but its redundant since you have to close the channel

	close(doneChannel)
}

func writeOnTerminal(doneChannel <-chan bool) { // doneChannel is a read-only channel (<-chan)
	for {
		select {
		case <-doneChannel: // Fufilled by closing the channel or passing a true value to it
			return
		default: // while the channel is open and has false as its value the loop keeps on going
			fmt.Println("Doing some work!")
		}
	}
}
