package main

import "fmt"

func main() {
	var slice []int

	fmt.Println(len(slice), cap(slice), slice) // 0 0 []

	if slice == nil {
		fmt.Println("nil!") // nil!
	}
}
