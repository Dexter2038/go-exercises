package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context, jobs chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task done, Im leaving")
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Job queue closed, I’m leaving")
				return
			}
			fmt.Println("Processing job ", job)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	jobs := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 5; i++ {
		go work(ctx, jobs)
	}
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)
	time.Sleep(10 * time.Second)
}
