package main

import (
	"context"
	"time"

)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				return
			default:
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
}

// This code demonstrates how to properly cancel a context in a goroutine
// without leaking resources. The goroutine checks for cancellation	

