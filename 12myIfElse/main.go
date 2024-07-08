package main

import "fmt"

func main() {
	fmt.Println("If Else in golang.")

	loginCourse := 23

	if loginCourse < 10 {
		fmt.Println("new user")
	} else if loginCourse < 20 {
		fmt.Println("regular user")
	} else {
		fmt.Println("old user")
	}
}
