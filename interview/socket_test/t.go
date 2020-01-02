package main

import "fmt"

func main() {
	m := make(map[string][]string)

	m["a"] = make([]string, 0)

	m["a"] = append(m["a"], "a")

	fmt.Printf("%T===%v", m["a"], m["a"])
}
