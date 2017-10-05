package main

import (
	"errors"
	"fmt"
)

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

	// Error types -----
	err := errors.New("emit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}

	// Define by group -----

	// Basic form.

	// import "fmt"
	// import "os"
	//
	// const i = 100
	// const pi = 3.1415
	// const prefix = "Go_"
	//
	// var i int
	// var pi float32
	// var prefix string

	// Group form.
	// seems better

	// import(
	//   "fmt"
	//   "os"
	// )
	//
	// const(
	//     i = 100
	//     pi = 3.1415
	//     prefix = "Go_"
	// )
	//
	// var(
	//     i int
	//     pi float32
	//     prefix string
	// )

	// enum for go style, you can declare even without iota.
	const (
		x = iota
		y = iota
		z = iota
		w
	)

	fmt.Printf("%d %d %d\n", x, y, z) // 0 1 2
	fmt.Printf("%d\n", w)             // 3

	const te = iota
	fmt.Printf("%d\n", te) // 0

	// interesting. when you keep bumping number. should multiple lines rather than single line
	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
	)
	fmt.Printf("%d %d %d\n", e, f, g) // 0

	// Any variable that begins with a capital letter means it will be exported, private otherwise.

	// array

	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	arr[9] = 1345
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	// a4 := [3]int{1, 2, 3} // define an int array with 3 elements

	b4 := [10]int{1, 2, 3}
	// define a int array with 10 elements, of which the first three are assigned.
	//The rest of them use the default value 0.
	fmt.Printf("int %d %d\n", b4[2], b4[9])

	c4 := [...]int{4, 5, 6} // use `…` to replace the length parameter and Go will calculate it for you.
	fmt.Printf("int %d %d\n", c4[1], c4[2])
}
