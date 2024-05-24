package main

import (
	"context"
	"fmt"
)

// context can have multiple values
func main() {

	key1 := "key1"
	key2 := "key2"

	ctx := context.Background()
	ctx = context.WithValue(ctx, key1, "some value")
	ctx = context.WithValue(ctx, key2, "some other value")

	fmt.Printf("value 1: %s\n", ctx.Value(key1))
	fmt.Printf("value 2: %s\n", ctx.Value(key2))
}
