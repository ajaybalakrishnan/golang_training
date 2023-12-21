package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome Files in go lang")
	content := "This needs to go in a file."

	file, err := os.Create(("./myfile.txt"))
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside teh file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
