package main

import "fmt"

func main() {
	fmt.Println("Structs in golang.")
	// no inheritance in golang; No Super or parent

	vamshi := User{"Vamshi", "vamshi@go.dev", true, 22}
	fmt.Println(vamshi)
	fmt.Printf("vamshi details are: %+v\n", vamshi)
	fmt.Printf("Name is %v and email is %v .", vamshi.Name, vamshi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
