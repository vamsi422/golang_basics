package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang.")

	//fruites := []string{"apple", "banana", "mango", "grapes", "orange"}

	// simple for loop

	// for i := 0; i < len(fruites); i++ {
	// 	fmt.Println(fruites[i])
	// }

	// for each loop using index

	// for index := range fruites {
	// 	fmt.Println(fruites[index])
	// }

	// for each loop using values

	// for _, fruit := range fruites {
	// 	fmt.Println(fruit)
	// }

	// while loop

	val := 1
	for val < 25 {
		if val == 22 {
			goto luck
		}
		fmt.Println(val)
		val++
	}

luck:
	fmt.Println("this is your luck number 22")
}
