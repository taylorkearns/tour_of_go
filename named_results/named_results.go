package main

import "fmt"

func doThings(x float64) (y, z float64) {
	y = x * 2
	z = x / 2

	return
}

func main() {
	fmt.Println(doThings(13))
}
