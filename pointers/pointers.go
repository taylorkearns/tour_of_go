package main

import "fmt"

func main() {
	i := 45

	p := &i

	*p = 56

	fmt.Println(*p)
}
