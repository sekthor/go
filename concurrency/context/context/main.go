package main

import (
	"context"
	"fmt"
)

func main() {
	var ctx context.Context

	// a completely empty context without deadline
	ctx = context.Background() // returns *context.emptyCtx -> implements context.Context
	Task(ctx)

	// TODO() is to be used as a placeholder, that will later be replaced by an actual context
	// use when implentenation of the context is still outstanding
	ctx = context.TODO() // returns *context.emptyCtx -> implements context.Context
	Task(ctx)

	// copies the parent context and adds the given value
	ctx = context.WithValue(ctx, "hello", "world") // returns *context.valueCtx
	Task(ctx)
}

func Task(ctx context.Context) {
	fmt.Printf("%T\n", ctx)
}
