package main

import "fmt"

// In golang there is no way to use ++variable,we need to use only variable++

func main() {
	fmt.Println("Welcome to Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("The slices are ", days)

	// Way 1 of using for loop
	for d := 0; d < len(days); d++ {
		fmt.Println("The value are ", days[d])
	}

	// Another way of using for loop
	for d := range days {
		fmt.Println("The slice data is ", days[d])
	}

	//Another way of using for loop
	for idx, val := range days {
		fmt.Printf("The index %v has %v\n", idx, val)
	}

	// Break statement
	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 5 {
			break
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

	// continue statement
	for rougeValue < 10 {

		if rougeValue == 5 {
			// we need to increment the variable else it causes an error
			rougeValue++
			continue
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

	// goto statement
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto vpk
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

vpk:
	{
		fmt.Println("Welcome to the world of golang")
	}

}
