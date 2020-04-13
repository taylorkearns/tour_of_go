package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type LongVertex struct {
	X, Y, Z int
}

var (
	v1 = LongVertex{1, 2, 3}
	v2 = LongVertex{Y: 2}
	v3 = LongVertex{}
	p1 = &LongVertex{4, 5, 6}
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v) // {1, 2}

	v.X = 4
	fmt.Println(v) // {4, 2}

	p := &v
	(*p).X = 6
	fmt.Println(v) // {6, 2}

	p.X = 8
	fmt.Println(v) // {8, 2}

	fmt.Println(v1)      // {1, 2, 3}
	fmt.Println(v2)      // {0, 2, 0}
	fmt.Println(v3)      // {0, 0, 0}
	fmt.Println(p1)      // &{4, 5, 6}
	fmt.Println((*p1).Z) // 6
	fmt.Println(p1.Z)    // 6
}
