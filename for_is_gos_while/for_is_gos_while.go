package main

import "fmt"

func main() {
	sum := 0

	for sum < 1000 {
		fmt.Println(sum)
		sum += 3
	}
}
