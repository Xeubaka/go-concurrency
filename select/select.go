package main

import (
	"fmt"
	"time"
)

/*
Select checks for the next channel with value and prints it

PERFORMANCE:
channel1 is waiting 0.5s
channel2 is waiting 2s

Select will prioritize channel1 while channel2 is still processing
*/
func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel 2"
		}
	}()

	for {
		select {
		case mensagemchannel1 := <-channel1:
			fmt.Println(mensagemchannel1)
		case mensagemchannel2 := <-channel2:
			fmt.Println(mensagemchannel2)
		}

	}
}
