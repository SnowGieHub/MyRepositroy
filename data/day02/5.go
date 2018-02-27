package main

import "fmt"

type Ainimal interface {
	Eat()
}

type Dog struct {
	yellow string
}

func (this *Dog) Eat() {

	fmt.Println("Dog eating ...")

}

type Cat struct {
	red string
}

func (this *Cat) Eat() {

	fmt.Println("Cat eating ...")

}

func test(ainimal Ainimal) {

	ainimal.Eat()

}

func main() {

	var dog Dog

	test(&dog)

}
