package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang")

	count := 20
	var result string

	if count < 20 {
		result = "The value is less then 20"
	} else if count > 20 {
		result = "The value is greater than 20"
	} else {
		result = "The value is equal to 20"
	}
	fmt.Println(result)

	if count == 20 {
		fmt.Println("The count is equal to 20")
	} else {
		fmt.Println("The count is not equal to 20")
	}

	// we can define a variable and initialize on the if condition
	if num := 3; num < 10 {
		fmt.Println("value is less than 3")
	} else {
		fmt.Println("The value is not less than 3")
	}

}
