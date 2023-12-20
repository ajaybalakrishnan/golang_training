package main

import "fmt"

func main() {
	fmt.Println("Inside main function")
	greeter()

	// greeterTwo()
	// func() { fmt.Println("Anonymous Function") }()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, str := proAdder(6, 7, 7, 8)
	fmt.Println("Result from pro adder is", proRes, "\n", str)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum = sum + val
	}
	return sum, "Pro result function"
}

func greeterTwo() {
	fmt.Println("Another Method")
}

func greeter() {
	fmt.Println("Hello! Welcome to functions in golang")
}
