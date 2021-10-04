package main

import "fmt"

// defer -> whenever defer key is added to a line it skips the line and run on the undiffered statement

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	// The out will be -> Hello \n World
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	// My output prediction -> Hello Three,Two,One World
	myDefer()
	// My prediction -> Hello 4,3,2,1,0,Three, Two ,One,World
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
