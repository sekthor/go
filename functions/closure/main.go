package main

import "fmt"

func closure() func(string) string {
	// this variable gets instanciated upon calling of closure()
	// closure() returns an anonymous child function
	// the child function can alwas reference variable "str" which is outside it's method scope
	str := ""
	return func(word string) string {
		str += word + " "
		return str
	}
}

func main() {
	// get the closure function
	concatter := closure()

	// call the closure function
	// every call of the closure will modify the original str variable
	fmt.Println(concatter("hello"))
	fmt.Println(concatter("beautiful"))
	fmt.Println(concatter("people"))
}
