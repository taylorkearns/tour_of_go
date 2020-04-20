package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	printSlice(slice) // len=5 cap=5 [1 2 3 4 5]

	slice = slice[:0]
	printSlice(slice) // []

	slice = slice[:4]
	printSlice(slice) // [1 2 3 4]

	p := &slice
	*p = slice[:0]
	printSlice(slice) // []

	slice = slice[:4]
	printSlice(slice) // [1 2 3 4]

	slice = slice[2:]
	printSlice(slice) // [3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
