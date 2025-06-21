package main

// Task:
// Context-Aware Worker Pool
//
// Write a Go program that:
//
// Spawns 3 worker goroutines.
// Each worker:
// Waits on a jobs channel`
// Prints the job ID (like Processing job 1)
// Sleeps 1 second to simulate work
// Listens for ctx.Done() to exit gracefully
// main():
// Sends 10 jobs into the channel
// Cancels everything after 5 seconds using context.WithTimeout

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
