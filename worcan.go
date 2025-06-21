package main

// Task:
// Worker Pool with Job Results and Cancellation
//
// Write a program that:
//
// Builds a worker pool like in task with workers (3 workers, job channel).
// Each job is an int (job ID).
// Workers process jobs by:
// Sleeping 1 second
// Returning a result: job * 2 (just double the job ID)
// Collect results in a separate results channel.
// The main goroutine must:
// Send 10 jobs
// Collect and print results as they come
// Cancel all work after 5 seconds using context.WithTimeout
// Stop receiving results after cancellation

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
