package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths using golang")

	var numberOne int = 5
	var numberTwo float64 = 6.5

	fmt.Println("The addition of two numbers is ", numberOne+int(numberTwo)) // we will loose precision here , if we add like this
	// without typecasrting the numberTwo the addition cannot be performed

	// Generate random number using Math
	//fmt.Println("The random number is ", rand.Intn(5))  // This will give us always a same number , to over come this we add seed to this

	// After adding seed
	// rand.Seed(55)
	// fmt.Println("The random number is ", rand.Intn(5)) // Again this will gives us same number as the seed is same for all the time ,
	// TO over come this we will pass a unique number to seed, which is time to the seed functio

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("The random number is ", rand.Intn(5))

	// Generate random number using crypto
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("The random number is ,:", randNumber)
}
