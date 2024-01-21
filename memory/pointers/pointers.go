package main

import "fmt"

func main() {
	var x int = 2
	fmt.Println(x)

	// use & to reference a value by it's memory adress
	ptr := &x

	// dereference the pointer to assgin a new vale
	*ptr = 5

	// x is changed since x and ptr point to the same memory adress
	fmt.Println(x)
}
