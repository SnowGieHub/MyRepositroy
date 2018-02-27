package main

//工厂模式

import "fmt"

type Game interface {
	show()
}

type Hero struct {
}

func (this *Hero) show() {

	fmt.Println("Hero show ...")

}

type Star struct {
}

func (this *Star) show() {

	fmt.Println("Star show")

}

func test(game Game) {

	game.show()

}

func main() {

	var hero Hero
	test(&hero)

	var star Star
	test(&star)

}
