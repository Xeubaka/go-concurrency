package main

import (
	"fmt"
	"time"
)

/*
	Generator creates a Channel with a encapsulated GoRoutine function
	generateChannel() returns a channel with a GoRoutine that fill its value every 0.5s
	since its not buffered it won't overflow, but it can be called in a loop.
*/

func main() {
	channel := generateChannel("Hello World!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func generateChannel(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
