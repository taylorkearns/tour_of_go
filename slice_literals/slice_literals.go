package main

import "fmt"

func main() {
	i := []int{1, 2, 3, 4, 5}
	b := []bool{true, true, false, false}
	s := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, false},
	}

	fmt.Println(i) // [1 2 3 4 5]
	fmt.Println(b) // [true true false false]
	fmt.Println(s) // [ {1 true} {2 true} {3 false} {4 false}]
}
