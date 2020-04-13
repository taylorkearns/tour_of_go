package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(j)

	p := &i
	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p)

	*p = *p / 37
	fmt.Println(*p)
	fmt.Println(j)
}

// 2701
// 42
// 21
// 2701
// 73
// 73
