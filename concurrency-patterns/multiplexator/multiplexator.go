package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Multiplexator will select from the next inputChannels with value and send it to the outputChannel
*/

func main() {
	channel := multiplexator(generateChannel("Hello World!"), generateChannel("Programming in GO!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexator(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case msg := <-inputChannel1:
				outputChannel <- msg
			case msg := <-inputChannel2:
				outputChannel <- msg
			}
		}
	}()

	return outputChannel
}

func generateChannel(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
