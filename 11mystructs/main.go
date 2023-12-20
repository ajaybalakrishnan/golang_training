package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")
	// no inheritance in golang; No super or parent

	ajay := User{"Ajay", "ajay@go.dev", true, 23}
	fmt.Println(ajay)

	fmt.Printf("Ajay Details are: %+v\n", ajay)
	fmt.Printf("Name is %v\t and email is %v\n", ajay.Name, ajay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
