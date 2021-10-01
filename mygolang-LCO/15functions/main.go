package main

import "fmt"

// The starting point for a go program to run is main function
// we cannot able to define a function with in a function in golang

func main() {
	fmt.Println("Welcome to functions in golang")

	greeter()

	result := adder(3, 5)
	fmt.Println("The result of addition is ", result)

	proResult := proAdder(1, 2, 3, 4, 5)
	fmt.Println("The result of proAdder is ", proResult)

	proResult, res := proAdderTwo(1, 2, 3, 4, 5)
	fmt.Println("The result of proAdder is ", proResult)
	fmt.Println("The result of proAdder is ", res)

}

func greeter() {
	fmt.Println("Hello world !")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(val ...int) int {
	tot := 0
	for _, d := range val {
		tot += d
	}
	return tot
}

func proAdderTwo(val ...int) (int, string) {
	tot := 0
	for _, d := range val {
		tot += d
	}
	return tot, "The proAdderTwo is completed"
}
