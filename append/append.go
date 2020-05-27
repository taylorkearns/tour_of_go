package main

import "tour_of_go/printing"

func main() {
	var s []int
	printing.PrintSlice(s)

	s = append(s, 0)
	printing.PrintSlice(s)

	s = append(s, 1)
	printing.PrintSlice(s)

	s = append(s, 2, 3, 4)
	printing.PrintSlice(s)
}
