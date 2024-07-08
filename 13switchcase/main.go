package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("give seed number for dice.")

	seed, _ := reader.ReadString('\n')

	fmt.Println("your input seed value is :", seed)

	num, _ := strconv.ParseInt(strings.TrimSpace(seed), 10, 32)

	rand.Seed(time.Now().UnixMicro() * num)
	val := rand.Intn(6) + 1

	switch val {
	case 1:
		fmt.Println("roll output", val)
	case 2:
		fmt.Println("roll output", val)
	case 3:
		fmt.Println("roll output", val)
	case 4:
		fmt.Println("roll output", val)
	case 5:
		fmt.Println("roll output", val)
	case 6:
		fmt.Println("roll output", val, "roll again")
	default:
		fmt.Println("Something went wrong!")

	}
}
