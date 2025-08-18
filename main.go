package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	// deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <- ctx.Done():
			break
		}
	}
	cancelCtx() // Cancel the context to signal doAnother to stop

	time.Sleep(100 * time.Millisecond) // Wait for doAnother to finish
	fmt.Printf("Doing something with context: %s\n", ctx.Value("key"))

	// anotherCtx := context.WithValue(ctx, "key", "anotherValue")
	// doAnother(anotherCtx)
	// fmt.Printf("Back in doSomething with context: %s\n", ctx.Value("key"))	
}

func doAnother(ctx context.Context, printCh <- chan int) {
	for {
		select {
			case <- ctx.Done():
				if err := ctx.Err(); err != nil {
					fmt.Printf("doAnother err: %s\n", err)
				}
				fmt.Printf("doAnother: finished\n")
				return
			case num := <-printCh:
				fmt.Printf("doAnother: received %d\n", num)
		}
	}
	fmt.Printf("Doing another thing with context: %s\n", ctx.Value("key"))
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "value")
	doSomething(ctx)
}