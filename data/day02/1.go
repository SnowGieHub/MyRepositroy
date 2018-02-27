package main

//函数指针  匿名函数
import "fmt"

func add(a int, b int) int {

	return a + b

}

func sub(a int, b int) int {

	return a - b
}

func calculator(a int, b int, op func(int, int) int) int {

	ret := op(a, b)
	return ret
}
func main() {

	ret := calculator(20, 10, add)
	fmt.Println("ret = ", ret)

	ret = calculator(200, 100, sub)
	fmt.Println("ret = ", ret)

	ret = func(x int, y int) int {

		return x + y

	}(1000, 19)
	fmt.Println("ret = ", ret)

	func() {
		fmt.Println("匿名函数调用")
	}()
}
