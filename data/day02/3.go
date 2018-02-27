package main

import "fmt"
import "MyPackAge"

func init() {

	fmt.Println("mian init")

}

func main() {

	fmt.Println(MyPackAge.Fun1(10, 20))

	fmt.Println(MyPackAge.Fun2(200, 100))

	var book MyPackAge.Book
	book.Name = "李雪楠"
	book.Age = 25

	fmt.Printf("% +v \n", book)

}
