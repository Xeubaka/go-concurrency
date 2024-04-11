package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string) // declaration of channel (unbuffered)

	go writeOnTerminal("Hello World!", channel)

	// for {
	/*
		Main will stop its process at the next line until channel has some value.
		<-channel returns both its value and its current state (open/close)
	*/
	// 	msg, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	/*
		Simpler solution, "range" will wait until channel has some value.
		Also it will check for its state, so if closed it end the for loop
	*/
	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("End of main()")
}

func writeOnTerminal(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto //Sends value to the channel
		time.Sleep(time.Second)
	}

	close(channel) //Always close the channel to prevent DEADLOCK
}
