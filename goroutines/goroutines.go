package main

import (
	"fmt"
	"time"
)

/*
	GoRoutines Concurrency:
	[Fork-Join] model

	[start] main()
		|
		|   [Forked]    [start] writeOnTerminal()
		|>>>>>>>>>>>>>>>>>|
		|                 |
		|                 |
		|                 |
		|   [Join]        |
		[X]<<<<<<<<<<<<<<<[X] // Join point on main()
		|                 |
		|                 |
	[end] main()        [end] writeOnTerminal()
*/

func main() {
	go writeOnTerminal("Hello World!")       //GoRoutines
	go writeOnTerminal("Programming in Go!") //GoRoutines

	time.Sleep(time.Second * 10)
}

func writeOnTerminal(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
