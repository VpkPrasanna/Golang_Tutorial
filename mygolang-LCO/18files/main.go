package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	constant := "This will go in to the files by vpk.in"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, constant)
	checkErrNil(err)
	fmt.Println("The length of the file is :", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	// reading a file in golang will return the readed data in bytes format
	dataBytes, err := ioutil.ReadFile(filename)
	checkErrNil(err)
	fmt.Println("The data present in the file is :", dataBytes)
	// To convert the byte data into string we use string wrapper around it
	fmt.Println("The data present in the file is :", string(dataBytes))

}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
