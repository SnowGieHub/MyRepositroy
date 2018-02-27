package main

import "fmt"
import "MyPackAge"

type Hero struct {
	Name string
	Age  int
}

func (this *Hero) SetName() {

	this.Name = "李雪楠"
	this.Age = 25

}

func (this *Hero) Show() {

	fmt.Println("name = ", this.Name)
	fmt.Println("age = ", this.Age)
}
func main() {

	var hero Hero
	hero.SetName()
	hero.Show()

}
