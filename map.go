package main

// Task:
// Map Initialization Panic
//
// Write a program that:
//
// Declares a map[string]int without initializing it.
// Tries to insert a key-value pair.
// Handles the panic gracefully (hint: use recover()).
// Bonus: Now fix the panic by properly initializing the map.

func main() {
	var m map[string]int

	defer func() {
		if recover() != nil {
			m = map[string]int{}
		}
	}()

	m["key"] = 0
}
