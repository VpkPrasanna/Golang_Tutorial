package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:300/learn?course=reactjs&paymentid=sfnksdjbfids"

func main() {
	fmt.Println("Welcome to urls in golang")

	fmt.Println("The url is :", myurl)

	//parsing the url
	res, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("The result is :", res.Scheme)
	fmt.Println("The result is :", res.Host)
	fmt.Println("The result is :", res.Path)
	fmt.Println("The result is :", res.Port())
	fmt.Println("The result is :", res.RawQuery)

	// this will return as a key-value pairs
	qparams := res.Query()

	fmt.Printf("The type of query params %T \n", qparams)

	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	// To Construct a url
	// we always need to pass the reference of the variable	 &
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "tuttcss",
		RawQuery: "user=vpk",
	}
	fmt.Printf("The type of other url %T\n", partsOfURL)
	oneMoreUrl := partsOfURL.String()
	fmt.Println("The other url:", oneMoreUrl)
}
