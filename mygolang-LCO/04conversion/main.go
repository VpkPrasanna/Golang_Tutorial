package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza Bot")
	fmt.Println("Give us a rating between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", inp)
	// the type will be string and we need to convert it into int or float based on our need
	fmt.Printf("The type of the input is %T,", inp)

	// This won't work as the input value will contain '\n' at the end and we need to remove it.
	// rating, err := strconv.ParseFloat(inp, 64)
	// To remove the \n at the end we use
	rating, err := strconv.ParseFloat(strings.TrimSpace(inp), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The converted value is ", rating+1)

	}
}
