package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	// create a context with a two second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel() // make sure we cancel the context, before we

	err := blockingLongTask(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel() // make sure we cancel the context, before we

	err = nonBlockingLongTask(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// this function takes five seconds and will error if context deadline exceeds
// however, it still blocks unitl the full five seconds have passed
func blockingLongTask(ctx context.Context) error {
	time.Sleep(5 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		return errors.New("context deadline exceeded")
	}
	return nil
}

// this tasks will return as soon as context deadline exceeded
// and not wait for the 5 secs to pass
func nonBlockingLongTask(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("context deadline exceeded")
	case <-time.After(5 * time.Second):
		return nil
	}
}
