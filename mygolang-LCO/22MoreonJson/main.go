package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Explore more on JSON in golang")
	EncodeJson()
}

// `` These are aliasis of the keys present in the struct, we can create in any name as we need
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - if we add - as a filed this particular key will not be shows as a part of result
	Tags     []string `json:"tags,omitempty"` // omit empty will not return the key where it is nil/null
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 299, "lco.dev", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "lco.dev", "bcd123", []string{"Fullstack", "js"}},
		{"Angular", 299, "lco.dev", "vpk123", nil},
	}
	// Packing the data into JSON
	response, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", response)
}
