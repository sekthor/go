package main

import (
	"context"
	"fmt"
)

func main() {

	// Background returns a pointer to an context.emptyContext
	// emptyContext implements context.Context with pointer receivers
	// therefore a pointer is returned
	ctx := context.Background()
	fmt.Printf("context: %i; pointer: %p\n", ctx, &ctx)

	// If we pass the context by value, we pass the pointer to the
	// previously created context.emptyContext.
	// So by passing the value, we copy the value of the pointer
	// in the new function scope, but not the actual emptyContext.
	// The value of emptyContext is not copied!
	// We only create a new pointer, that points to the same emptyContext
	PassByValue(ctx)

	// so if we pass by reference, we create a pointer to our existing pointer
	// ptr -> ptr -> context.emptyContext
	PassByReference(&ctx)
}

func PassByValue(ctx context.Context) {
	fmt.Printf("context: %i; pointer: %p\n", ctx, &ctx)
}

// ctx is a new points to the existing pointer passed into the function
func PassByReference(ctx *context.Context) {
	fmt.Printf("context: %i; pointer: %p\n", *ctx, ctx)
}
