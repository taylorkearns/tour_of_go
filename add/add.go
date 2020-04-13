package main

import "fmt"

func add(n int, m int) int {
	return n + m
}

func main() {
	const a, b int = 3, 5
	sum := func(int, int) int { return a + b }

	fmt.Println(sum(a, b))
}
