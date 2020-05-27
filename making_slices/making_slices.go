package main

import printing "tour_of_go/printing"

func main() {
	// make and print slice w/ len 5 and cap 5
	a := make([]int, 5)
	printing.PrintNamedSlice("a", a)

	// make and print slice w/ len 0 and cap 5
	b := make([]int, 0, 5)
	printing.PrintNamedSlice("b", b)

	// make and print slice ending at 2
	c := b[:2]
	printing.PrintNamedSlice("c", c)

	// make and print slice starting at 2 and ending at 5
	d := c[2:5]
	printing.PrintNamedSlice("d", d)
}
