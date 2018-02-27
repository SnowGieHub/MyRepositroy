package main

import "fmt"

var aa, bb = 10, 20

func fun(a *int, b *int) {

	var temp int

	temp = *a
	*a = *b
	*b = temp

}

func main() {

	fun(&aa, &bb)

	fmt.Println("aa = ", aa)
	fmt.Println("bb = ", bb)
}
