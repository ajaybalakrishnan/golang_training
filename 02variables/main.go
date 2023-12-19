package main

import "fmt"

const LoginToken string = "ajklsfhiawhrfgi" // Public constant

func main() {
	var username string = "ajay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable if of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable if of type: %T \n", smallVal)

	var smallFloat float64 = 255.3345452453523452433435
	fmt.Println(smallFloat)
	fmt.Printf("Variable if of type: %T \n", smallFloat)

	// defaiult values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable if of type: %T \n", anotherVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable if of type: %T \n", anotherStringVariable)

	// impicit type

	var website = "ajaybalakrishnan.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable if of type: %T \n", LoginToken)
}
