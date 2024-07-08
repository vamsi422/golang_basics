package main

import "fmt"

// jwtTokem:= 3000 (not allowed)
var jwtTokem int = 3000

const LoginToken string = "ghabbhhjd" // public

func main() {
	var userName string = "vamshi"
	fmt.Println(userName)
	fmt.Printf("Variable is of type : %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.45372882382892898289389
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "google.com"
	fmt.Println(website) // cannot change website to int or any other type.

	// no var style

	numberOfUsers := 30000.0 // can be changed to any data type.
	fmt.Println(numberOfUsers)

	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
