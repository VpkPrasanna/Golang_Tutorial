package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	// The result will be like this The values in the array is ,: [Apple Tomato  Peach],
	// where the space between tomato and peach is 2 since it has alloted memory for 3 rd postition
	fmt.Println("The values in the array is ,:", fruitList)

	// This will give us the answer as 4 as we have initialized with 4 as the total input and it will create a an array of 4 inp with memory
	fmt.Println("The len of the fruit list is:,", len(fruitList))

	// Another way to define the sttings

	var vegList = [5]string{"Potato", "tomato", "beans"}
	fmt.Printf("The type of vegList is %T\n", vegList)
	fmt.Println("The Veglist is :,", vegList)
	fmt.Println("The len of the veg list is :", len(vegList))

}
