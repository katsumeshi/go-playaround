package main

import "fmt"

func main() {

	x := 1
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// initialize x, then check if x greater than
	if g := 5; g > 10 {
		fmt.Println("g is greater than 10")
	} else if g == 5 {
		fmt.Println("g is equal 5")
	} else {
		fmt.Println("g is less than 10")
	}

	// the following code will not compile
	//fmt.Println(g)
}
