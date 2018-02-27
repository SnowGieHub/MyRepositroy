package main

import "fmt"

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {

	fmt.Println("a = ", a, "b = ", b, "c = ", c)

	fmt.Println("d = ", d, "e = ", e, "f = ", f)

	fmt.Println("g = ", g, "k = ", k)

}
