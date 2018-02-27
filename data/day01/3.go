package main

import "fmt"

var aa, bb = 25, "lixuenan"

func boo(a int, b string) int {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 404

}
func main() {

	var aaa int = boo(aa, bb)

}
