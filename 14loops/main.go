package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 4 {
			goto flag
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		} else if rougueValue == 6 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

flag:
	fmt.Println("Flag is true")
}
