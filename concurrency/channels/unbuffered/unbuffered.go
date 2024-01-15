package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// channels allow sending and receiving of data

	// create a channel of type integer
	ch := make(chan int)
	go Task(ch)

	// receiving from channel: blocks until data is received
	var x int = <-ch

	fmt.Println(x)
}

func Task(c chan int) {
	// artificial delay -> simulate long task
	time.Sleep(time.Second)

	// send an integer into the channel
	c <- rand.Int()
}
