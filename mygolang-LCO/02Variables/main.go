package main

import "fmt"

var jwtToekn = 500

const loginToken string = "vnkdsnvksdn" //public variable

func main() {
	var userName string = "Prasanna Kumar V"
	fmt.Println(userName)
	fmt.Printf(" Variable of type : %T \n", userName)

	var isPresent bool = false
	fmt.Println(isPresent)
	fmt.Printf(" Variable of type : %T \n", isPresent)

	var id uint32 = 256
	fmt.Println(id)
	fmt.Printf(" Variable of type : %T \n", id)

	var id1 float32 = 256.0789456123
	fmt.Println(id1)
	fmt.Printf(" Variable of type : %T \n", id1)

	//default values
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf(" Variable of type : %T \n", anothervariable)

	var emptyString string
	fmt.Println(emptyString)
	fmt.Printf(" Variable of type : %T \n", emptyString)

	//implicit type
	var oneMorevariable = "Prasanna"
	fmt.Println(oneMorevariable)
	fmt.Printf(" Variable of type : %T \n", oneMorevariable)

	// no var style of initialization
	//temp := 8
	fmt.Println(jwtToekn)
	fmt.Printf(" Variable of type : %T \n", jwtToekn)

	fmt.Println(loginToken)
	fmt.Printf(" Variable of type : %T \n", loginToken)
}
