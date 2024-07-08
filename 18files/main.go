package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang.")
	content := "this is the content in this file."
	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
