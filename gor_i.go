package main

// Task:
// Goroutine Capture Trap
//
// Write a program that:
//
// Starts five goroutines inside a loop.
// Each one should print its own index (i) — not the final value of i.
// Wait for all of them to finish.
// The trap: don’t fall for the infamous closure bug.

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
