package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		// We can get the error from the context
		err := context.Cause(ctx)
		fmt.Println("halted operation2 due to error: ", err)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)

	go func() {
		err := operation1(ctx)
		if err != nil {
			// this time, we pass in the error as an argument
			cancel(err)
		}
	}()

	// Run operation2 with the same context we use for operation1
	operation2(ctx)
}
