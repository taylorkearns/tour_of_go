package main

import "fmt"

func myPic(dx, dy int) [][]uint8 {
	result := [][]uint8{}

	for i := 0; i < dy; i++ {
		result = append(result, childSlice(i, dx))
	}

	return result
}

func childSlice(parentI, length int) []uint8 {
	result := []uint8{}

	for i := 0; i < length; i++ {
		result = append(result, uint8(parentI*i))
	}

	return result
}

func main() {
	fmt.Println(myPic(3, 5))
}
