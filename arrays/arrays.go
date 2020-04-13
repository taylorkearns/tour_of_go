package main

import "fmt"

var beatles = [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}

func main() {
	var a [2]string
	a[0] = "Robin"
	a[1] = "Charlotte"
	fmt.Println(a[0], a[1]) // Robin Charlotte
	fmt.Println(a)          // [Robin Charlotte]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]

	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]

	s1 := primes[0:5]
	fmt.Println(s1) // [2 3 5 7 11]

	fmt.Println(beatles)

	jp := beatles[0:2]
	gr := beatles[2:4]
	fmt.Println(jp, gr) // [John Paul] [George Ringo]

	jp[1] = "Jim"
	gr[0] = "Jude"
	fmt.Println(jp, gr)  // [John Jim] [Jude Ringo]
	fmt.Println(beatles) // [John Jim Jude Ringo]
}
