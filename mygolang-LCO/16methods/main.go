package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

func main() {
	fmt.Println("Welcome to Methods in golang")

	// call the defined struct
	vpk := User{"Prasanna Kumar", "vpk@gmail.com", true, 17}
	fmt.Println("The Struct contains :", vpk)
	// %+v in struct gives us more details in key and value format defined in struct
	fmt.Printf("The struct values are %+v\n", vpk)
	fmt.Printf("Prasanna email is %v and age is %v\n", vpk.Email, vpk.Age)
	vpk.GetStatus()
	// we are calling the function of changed email
	vpk.NewMail()
	// when again we are calling the function it prints the old value , which means a copy of object is passed as an input . To overcome that we need pass the reference of the object.
	fmt.Printf("Prasanna email is %v and age is %v\n", vpk.Email, vpk.Age)

}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Verified)
}

// Here we are changing the value of the email with an object
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email if this user is :", u.Email)
}
