package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "John Doe", "Your name (use -name <your_name>)")
	flag.Parse() // if we put defer before this statement the output will be default value.
	fmt.Printf("Hello, %s!\n", name)
}
