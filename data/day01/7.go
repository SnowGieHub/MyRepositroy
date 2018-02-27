package main

import "fmt"

func main() {

	defer fmt.Println("end")
	defer fmt.Println("end2")

	fmt.Println("1111")
	fmt.Println("2222")
	fmt.Println("3333")

}
