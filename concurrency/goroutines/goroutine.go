package main

import "fmt"

func main() {
	// the goroutine for Task() is started, but without synchronization
	// the main method will exit while Task() is still running
	// effectively aborting execution of the goroutine
	go Task()
	fmt.Println("main thread")
}

func Task() {
	fmt.Println("start task")
	fmt.Println("end task")
}
