package main

// Task:
// Context Cancellation with select
//
// Write a program that:
//
// Starts a goroutine that prints a message every second.
// Cancels the goroutine after 3 seconds using context.WithCancel.
// Uses select inside the goroutine to listen for the cancel signal.

import (
	"context"
	"fmt"
	"time"
)

func message(ctx context.Context) {
	fmt.Println(time.Now())
	fmt.Println("message")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Bye")
			return
		case <-ticker.C:
			fmt.Println(time.Now())
			fmt.Println("message")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go message(ctx)
	time.Sleep(4 * time.Second)
}
