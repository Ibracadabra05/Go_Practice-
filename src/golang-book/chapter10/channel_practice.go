/* CHANNELS */

package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) { //can only send to this channel
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) { //can only send to this channel
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { //can only receive from this channel
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
