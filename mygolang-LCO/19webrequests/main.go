package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests in golang")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// The type of the response will be a reference of the response we received
	fmt.Printf("The type of response is %T \n", res)

	// Close the connection
	defer res.Body.Close()

	// To read the content present in the url
	dataBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The response contains ", string(dataBytes))
}
