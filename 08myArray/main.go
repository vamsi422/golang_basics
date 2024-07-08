package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[3] = "Tomato"
	fmt.Println("fruits array values are :", fruitList)
	fmt.Println("fruits array length is :", len(fruitList))

	var vegList = [3]string{"onion", "lemon", "cucumber"}
	fmt.Println("vegy array values are :", vegList)
	fmt.Println("fruits array length is :", len(vegList))

}
