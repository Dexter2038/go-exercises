package main

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
