package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "UserID", 123)

	var wg sync.WaitGroup
	wg.Add(1)
	go Task(ctx, &wg)
	wg.Wait()
	processRequest(ctx)
}

func Task(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	userId := ctx.Value("UserID")
	fmt.Println("User ID is", userId)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("UserID").(int)
	fmt.Println("Processing request for User ID:", userID)
}