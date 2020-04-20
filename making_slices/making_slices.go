package main

import "fmt"

func main() {
	// make and print slice w/ len 5 and cap 5
	a := make([]int, 5)
	printSlice("a", a)

	// make and print slice w/ len 0 and cap 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	// make and print slice ending at 2
	c := b[:2]
	printSlice("c", c)

	// make and print slice starting at 2 and ending at 5
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(sliceName string, slice []int) {
	fmt.Printf(
		"%s, len=%d, cap=%d, %v\n",
		sliceName, len(slice), cap(slice), slice,
	)
}
