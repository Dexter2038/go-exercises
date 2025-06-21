package main

// Task:
// Buffered Channels
//
// Write a program that:
//
// Creates a buffered channel of type chan string with a capacity of 3.
// Starts three goroutines, each one sending a letter ("A", "B", "C") into the channel.
// In main(), reads and prints all three values from the channel.

import "fmt"

func add(ch chan string, s string) {
	ch <- s
}

func main() {
	ch := make(chan string, 3)
	go add(ch, "A")
	go add(ch, "B")
	go add(ch, "C")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
