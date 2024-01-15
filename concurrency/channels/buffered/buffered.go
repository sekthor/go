package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// int channel has a buffer of size 2
	ch := make(chan int, 2)

	go Task(ch)

	// receiving will block if the buffer is emtpy
	fmt.Println(<-ch)

	ch <- 1
	go Task(ch)

	// but not when it has data
	fmt.Println(<-ch) // 1
	// emtpty the buffer
	fmt.Println(<-ch) // math/rand

	// fill the buffer
	ch <- 1
	ch <- 2
	go Task(ch)

	time.Sleep(time.Second * 2) // go routine is blocked until these two seconds pass, as the buffer is full
	fmt.Println(<-ch)           // 1
	fmt.Println(<-ch)           // 2
	fmt.Println(<-ch)           // math/rand from goroutine

	time.Sleep(time.Second) // naive wait for goroutine to finish
}

func Task(ch chan int) {
	start := time.Now()
	time.Sleep(time.Second)
	ch <- rand.Int()
	elapsed := time.Since(start)
	fmt.Printf("goroutine took %fs\n", elapsed.Seconds())
}
