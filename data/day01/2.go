package main

import "fmt"

func foo(a int, b string) (ret1 int, ret2 string) {

	ret1 = 25
	ret2 = "李雪楠"

	return

}

func main() {

	ret1, ret2 := foo(100, "lixuen")

	fmt.Println("姓名：", ret2, "年龄：", ret1)

}
