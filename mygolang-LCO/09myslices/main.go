package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("The type of vegList is %T\n", fruitList)

	//To add more values to the defined slices
	fruitList = append(fruitList, "Orange", "banana")
	fmt.Println("The updated fruit list is ", fruitList)

	// Selecting subset of slices, which will print from index 1 to end of slice
	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)
	// This will print from 2 to 4 , which will leave the value of 5
	// fruitList = append(fruitList[2:5])
	// fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	//highscores[4] = 777

	// we can use append method to add more values to our slices than it was initialized!
	highscores = append(highscores, 555, 666, 321)
	fmt.Println(highscores)

	// Sorting in slices
	sort.Ints(highscores)
	fmt.Println(highscores)

	// To check whether slice is sorted or not
	value := sort.IntsAreSorted(highscores)
	fmt.Println(value)

}
