package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// no inheritance in golang ,no super or parent
	vamshi := User{"vamshi", "me@gmail.com", true, 23}
	fmt.Println(vamshi)
	fmt.Println("User status is :", userStatus(vamshi))

	fmt.Println("enter your new email:")
	reader := bufio.NewReader(os.Stdin)

	newEmail, _ := reader.ReadString('\n')

	vamshi = setEmail(vamshi, newEmail)

	fmt.Println("your input is :", vamshi.Email)

}

type User struct {
	Name     string
	Email    string
	isActive bool
	Age      int
}

func userStatus(u User) bool {
	return u.isActive
}
func setEmail(u User, newEmail string) User {
	u.Email = newEmail
	return u
}
