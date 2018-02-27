package main

//类型断言

import "fmt"

func test(arry interface{}) {

	value, ok := arry.(string)
	if ok {
		fmt.Println("string 类型", "value = ", value)
	} else {
		fmt.Println("不是string 类型")
	}

}

func main() {

	a := 25
	b := "李雪楠"

	test(a)
	test(b)
}
