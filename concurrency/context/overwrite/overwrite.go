package main

import (
	"context"
	"fmt"
)

// we can change a value of a context
func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "value1")
	fmt.Printf("value: %s\n", ctx.Value("key"))

	ctx = context.WithValue(ctx, "key", "value2")
	fmt.Printf("value: %s\n", ctx.Value("key"))
}
