package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go Tasks(ctx)
	time.Sleep(3 * time.Second)

}

func Tasks(ctx context.Context) {
	select {
	case <- ctx.Done():
		fmt.Println("Deadline exceeded", ctx.Err())
		return
	}
}