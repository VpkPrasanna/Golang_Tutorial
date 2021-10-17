package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request in golang ")
	// Function call for GET request
	// performGetRequest()

	// Function call for POST request
	perfromJsonPostRequest()
}

func performGetRequest() {
	const myUrl string = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// Another way to print the content from the request using builder function from strings package.
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println("The content in the request is :", responseString.String())

	// Regular way to print the content as output
	// fmt.Println("The content in the url is :", string(content))
}

func perfromJsonPostRequest() {
	const myurl string = "http://localhost:8000/post"

	//create fake json
	fakedata := strings.NewReader(`
	{
		"Name":"VPK"
	}`)

	res, err := http.Post(myurl, "application/json", fakedata)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	contentByte, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contentByte))
}
