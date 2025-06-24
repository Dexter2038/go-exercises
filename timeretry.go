package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task:
// Per-Job Timeout & Retry
//
// Write a program that:
//
// Spawns 3 workers from a pool.
// Receives 10 jobs, each being an integer.
// Each job:
//   Has a random processing time between 1–4 seconds.
//   Must complete within 2 seconds or it times out.
//   If it times out, it gets retried up to 2 times total (1 retry).
// Collects and prints the successful results.
// Cancels all remaining work after 15 seconds total (context.WithTimeout on the app).

func process_job(ctx context.Context, job int) error {
	duration := time.Duration(rand.Intn(4)+1) * time.Second
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	select {
	case <-time.After(duration):
		fmt.Printf("Job %d completed in successful way\n", job)
		return nil
	case <-ctx.Done():
		fmt.Printf("Job %d timed out after %v\n", job, duration)
		return ctx.Err()

	}
}

func work(id int, ctx context.Context, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutdown by global timeout\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d find out that jobs channel is closed\n", id)
				return
			}
			for attempt := 0; attempt < 2; attempt++ {
				err := process_job(ctx, job)
				if err == nil {
					results <- job
					break
				}
				fmt.Printf("Worker %d retries to run job %d in %d attempt\n", id, job, attempt+1)
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(i, ctx, jobs, results, &wg)
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
			fmt.Println("Main program died due to a timeout")
			return
		case result, ok := <-results:
			if !ok {
				fmt.Println("Main program find out that results channel is closed")
				return
			}
			fmt.Printf("Result %d is executed successfully\n", result)
		}
	}
}
