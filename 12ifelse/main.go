package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in golang")

	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// loginCount, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	loginCount := 25

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out!"
	} else {
		result = "Exactly 10 loging"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := loginCount; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {
	// 	statement
	// }
}
