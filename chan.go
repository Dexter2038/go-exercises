package main

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
