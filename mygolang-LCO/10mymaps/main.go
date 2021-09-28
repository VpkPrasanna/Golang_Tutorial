package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")

	// To define a map
	courseMap := make(map[string]string)

	courseMap["JS"] = "javaScript"
	courseMap["PY"] = "python"
	courseMap["RB"] = "Ruby"
	fmt.Println("The map values are,:", courseMap)

	// TO get value based on key
	fmt.Println("The full form of PY is :", courseMap["PY"])

	// TO delete a key
	delete(courseMap, "RB")

	fmt.Println("The Updated CourseMap is ", courseMap)

	// For loop in map
	for key, value := range courseMap {
		fmt.Printf("The key is %v and the value is %v\n", key, value)
	}

}
