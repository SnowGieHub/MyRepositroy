package main

import "fmt"
import _ "reflect"

func swap(arry []int) {

	i, j, temp := 0, 0, 0

	for i = 0; i < len(arry); i++ {

		for j = 0; j < len(arry)-i-1; j++ {

			if arry[j] > arry[j+1] {

				temp = arry[j]
				arry[j] = arry[j+1]
				arry[j+1] = temp

			}
		}

	}

}

func swap_1(arry []int) {

	i, j, temp := 0, 0, 0

	for i = 0; i < len(arry); i++ {

		for j = i + 1; j < len(arry); j++ {

			if arry[j] < arry[i] {
				temp = arry[j]
				arry[j] = arry[i]
				arry[i] = temp
			}
		}

	}

}

func swap_2(arry []int) {

	i, j, temp := 0, 0, 0

	for i = 1; i < len(arry); i++ {
		if arry[i] < arry[i-1] {

			temp = arry[i]

			for j = i - 1; j >= 0 && temp > arry[j]; j-- {

				arry[j+1] = arry[j]

			}
			arry[j+1] = temp

		}
	}

}

func swap_3(arry []int, start int, end int) {

	i := start
	j := end
	target := arry[start]
	if i < j {

		for i < j {

			for i < j && arry[j] < target {

				j--
			}
			if i < j {
				arry[i] = arry[j]
				i++
			}

			for i < j && arry[i] > target {
				i++
			}
			if i < j {
				arry[j] = arry[i]
				j--
			}

		}
		arry[i] = target

		swap_3(arry, start, i-1)

		swap_3(arry, i+1, end)

	}
}

func main() {

	arry := []int{10, 34, 56, 57, 67, 78, 9, 54, 543, 43}
	len := len(arry)
	swap_3(arry, 0, len-1)

	for _, value := range arry {

		fmt.Println("value = ", value)

	}

}
