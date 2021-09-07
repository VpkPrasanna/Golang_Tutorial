package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my Pizza Bot"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for PIZZA")

	// comma ok || comma error which is equal to try catch block
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", input)
	fmt.Printf("The type of variable %T", input)
}
