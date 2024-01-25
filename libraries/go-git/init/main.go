package main

import (
	"log"

	"github.com/go-git/go-git/v5"
)

func main() {
    _, err := git.PlainInit("./test.git", true)

	if err != nil {
		log.Fatal(err)
	}
}
