package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Package ")

	presentTime := time.Now()
	// This will give us the result in ugly format like -> The time is  2021-09-14 23:11:08.2579918 +0530 IST m=+0.004796901
	fmt.Println("The time is ", presentTime)
	// We will format this , to see in a better way like . : We need to follow this syntax to format the date
	updatedTIme := presentTime.Format("01-02-2006 15:04:05 Monday ")
	fmt.Println("The prettier version ,", updatedTIme)

	//Lets create a date
	newDate := time.Date(2020, time.July, 1, 23, 23, 0, 0, time.UTC)
	fmt.Println("The created date is ,", newDate)
	fmt.Println("The formatted date is ,", newDate.Format("01-02-2006 Monday"))

	// This part of code is from runtime packge of go-lang
	fmt.Println("The total number of cpu are ,", runtime.NumCPU())
}
