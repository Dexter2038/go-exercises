package main

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
