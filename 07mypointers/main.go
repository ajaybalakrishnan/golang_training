package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int

	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of the actual pointer is ", ptr)
	fmt.Println("Value of the actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Valuse is: ", myNumber)
}
