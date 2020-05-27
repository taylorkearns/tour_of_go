package printing

import "fmt"

func PrintSlice(slice []int) {
	fmt.Printf(
		"len=%d, cap=%d, %v\n",
		len(slice), cap(slice), slice,
	)
}

func PrintNamedSlice(sliceName string, slice []int) {
	fmt.Printf(
		"%s, len=%d, cap=%d, %v\n",
		sliceName, len(slice), cap(slice), slice,
	)
}
