package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context, jobs chan int, results chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Im done")
			return
		case job := <-jobs:
			results <- job * 2
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for i := range 3 {
	}
}
