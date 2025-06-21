package main

func main() {
	var m map[string]int

	defer func() {
		if recover() != nil {
			m = map[string]int{}
		}
	}()

	m["key"] = 0
}
