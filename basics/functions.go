package main

import (
	"fmt"
)

func main() {
	superman()
	result := multiply(2, 4)
	fmt.Println(result)
	results, mylen, myname := adder(3, 6, 8, 3, 45)
	fmt.Println(results)
	fmt.Println(mylen)
	fmt.Println(myname)
	k1, val := multiplier(5, 6, 8)
	fmt.Println("calling multplier function")
	fmt.Println(k1)
	fmt.Println(val)
}

func superman() {
	fmt.Println("I am super man")
}

func multiply(v1, v2 float64) float64 {
	return v1 * v2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	length := len(values)
	name := "Prasanna "
	return sum, length, name
}

func multiplier(datas ...int) (int, string) {
	total := 1
	for i := range datas {
		total = total * datas[i]
	}
	names := "lco"
	return total, names
}
