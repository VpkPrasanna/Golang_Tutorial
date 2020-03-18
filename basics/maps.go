package main

import (
	"fmt"
)

func main() {
	// make new
	// new  - only allocates - no init of memory
	// make  - allocate and init - non zeored storage

	score := make(map[string]int)
	score["hitesh"] = 89
	fmt.Println(score)

	score["josh"] = 34
	score["ron"] = 23
	score["sam"] = 56
	score["vicky"] = 78
	fmt.Println(score)

	getRonScore := score["ron"]
	fmt.Println(getRonScore)

	delete(score, "vicky")

	fmt.Println(score)

	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}

}
