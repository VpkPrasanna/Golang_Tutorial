package main

import "fmt"

// Struct Name has to start with capital Letter and all the attributes of the struct has to start with captial letter
type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

func main() {
	fmt.Println("Welcome to structs in golang")
	// No inheritance , NO classes , NO super Method

	// call the defined struct
	vpk := User{"Prasanna Kumar", "vpk@gmail.com", true, 17}
	fmt.Println("The Struct contains :", vpk)
	// %+v in struct gives us more details in key and value format defined in struct
	fmt.Printf("The struct values are %+v\n", vpk)
	fmt.Printf("Prasanna email is %v and age is %v\n", vpk.Email, vpk.Age)

}
