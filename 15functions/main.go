package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in go.")
	greeter()
	fmt.Println("output of adder function is:", adder(3, 12))
	sum, message := proAdder(2, 3, 4, 12, 5, 2, 9, 8)
	fmt.Println("proAdder result is :", sum, "\n", message)

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "sum calculated successfully."

}
func greeter() {
	fmt.Println("Hello ,nice to meet you!.")
}
func adder(a int, b int) int {
	return a + b
}
