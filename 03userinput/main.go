package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give feedback points for our pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type of The rating is  %T .", input)
}
