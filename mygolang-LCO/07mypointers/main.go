package main

import "fmt"

func main() {
	fmt.Println("Welcome to the lesson on Pointers")

	// Initialize the pointer and see what it has
	var ptr *int
	fmt.Println("The value inside the pointer is,", ptr)

	myNumber := 23
	// & operator is used to store the memory address of the variable
	var myNumberReference = &myNumber

	fmt.Println("The actual value of pointer is ,", myNumberReference)
	// Checking the type op refernce object
	fmt.Printf("The type of reference pointer is %T \n,", myNumberReference)

	// * operator is used to get the value present in the memory address
	fmt.Println("The actual value of pointer is ,", *myNumberReference)

	*myNumberReference = *myNumberReference + 3
	fmt.Println("The updated value is:", myNumber)
}
