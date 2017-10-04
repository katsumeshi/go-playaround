package main

import "fmt"

var enabled, disabled = true, false // global

func main() {

	// Define variables -----

	// var vname1, vname2, vname3 string
	var vname4, vname5, vname6 string = "a", "b", "c"
	fmt.Printf("%s%s%s\n", vname4, vname5, vname6) // abc

	vname7 := "a"
	fmt.Printf("%s%s%s\n", vname7) // a%!s(MISSING)%!s(MISSING)%

	vname8, vname9 := "a", "b"
	fmt.Printf("%s%s%s\n", vname8, vname9) // a%!s(MISSING)%!s(MISSING)%

	_, b := 34, 35      // _ : any value that is given to it will be ignored.
	fmt.Printf("%d", b) // 35

	// Constants -----

	const Pi = 3.14
	const Pipi = Pi
	const Ottupi = Pipi
	const prefix = "tikubi"
	fmt.Printf("%f\n", Pi)     // 353.140000????
	fmt.Printf("%v\n", prefix) // tikubi%

	// Boolean -----

	fmt.Printf("%v\n", enabled) // tikubi%

	// var enabled = false can't assign to global variable

	// Numerical types
	// can't compute diffrent integer type
	// var a int8 = 1
	// var b int32 = 2
	// c := a + b

	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", c)

	// String -----

	japaneseHello := "Ohaiou" // WTF ohayou is correct lol
	fmt.Printf("Value is: %v\n", japaneseHello)

	// var s string = "hello"
	// s[0] = 'c'

	s := "hello"
	as := []byte(s) // convert string to []byte type
	as[0] = 'c'
	s2 := string(as) // convert back to string type
	fmt.Printf("%s\n", s2)

	s1 := "hello,"
	m1 := " world"
	a1 := s1 + m1
	fmt.Printf("%s\n", a1)

	s3 := "hello"
	s3 = "c" + s3[1:] // you cannot change string values by index, but you can get values instead.
	fmt.Printf("%s\n", s3)

	m := `hello
    world`
	fmt.Printf("%s\n", m) // ` will not escape any characters in a string.

}
