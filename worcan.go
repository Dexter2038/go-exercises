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
	"sync"
	"time"
)

func work(ctx context.Context, jobs chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Im done")
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Jobs channel is closed")
				return
			}
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				return
			case results <- job * 2:
			}
		}
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(ctx, jobs, results, &wg)
	}

	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelling work")
			return
		case result, ok := <-results:
			if !ok {
				fmt.Println("results channel is closed")
				return
			}
			fmt.Println(result)
		}
	}
}
