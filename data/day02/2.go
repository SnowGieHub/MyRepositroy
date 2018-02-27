package main

import "fmt"

type Book struct {
	name string

	age int
}

func bk(val *Book) {
	val.name = "李"

}

func main() {

	var book Book
	book.name = "李雪楠"
	book.age = 25
	bk(&book)
	fmt.Printf("%+v \n ", book)

}
