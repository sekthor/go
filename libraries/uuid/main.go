package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// create new uuid
	id := uuid.New()

	fmt.Println(id)

	id, err := uuid.Parse("bbb82b2d-3a40-4c57-b268-a5f02120be47")

	if err != nil {
		fmt.Println("couldn not parse uuid")
	}

	fmt.Println(id)

	// the same as: uuid.Must(uuid.Parse(""))
	// will panic
	uuid.MustParse("")

}
