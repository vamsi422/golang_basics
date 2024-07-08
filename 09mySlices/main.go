package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"} // in slice we do not specify length of array
	fmt.Printf("Type of fruitList is %T. \n", fruitList)
	fruitList = append(fruitList, "Mangoes", "Banana")

	fmt.Println(fruitList)
	var ind int = 1
	fruitList = append(fruitList[:ind], fruitList[ind+1:]...)
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 432
	highScores[1] = 124
	highScores[2] = 613
	highScores[3] = 751
	//highScores[4] = 356 // throws error
	highScores = append(highScores, 344, 333, 212)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
