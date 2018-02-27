package main

import "fmt"

func main() {

	score := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(score); i++ {
		switch {
		case score[i] == 30:
			fmt.Println(score[i])
		}

	}

	/*
		for index, value := range score {

			fmt.Println("index = ", index)
			fmt.Println("value = ", value)

		}
	*/

}
