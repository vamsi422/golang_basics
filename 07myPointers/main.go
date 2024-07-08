package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int

	val := 12
	ptr = &val // providing reference to the ponter
	valTwo := 15
	var dirVal = &valTwo
	fmt.Println("pointer memory address is :", dirVal)
	fmt.Println("pointer value stored in the memory location is :", *dirVal)

	fmt.Println("pointer memory address is :", ptr)
	fmt.Println("pointer value stored in the memory location is :", *ptr)

}
