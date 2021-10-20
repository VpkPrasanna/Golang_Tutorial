package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Explore more on JSON in golang")
	// EncodeJson()

	DecodeJson()
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
	// MarshallIndent will Indent the json data with the specified data format
	response, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", response)
}

// Decode the json
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"Price": 299,
		"website": "lco.dev",
		"tags": ["web-dev","js"]
	}
	
	`)
	var lcoCourse course
	validJson := json.Valid(jsonDataFromWeb)
	if validJson {
		fmt.Println("The json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("IT IS NOT A VALID JSON")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v and the Type of value is %T \n", k, v, v)
	}

}
