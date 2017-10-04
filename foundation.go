package main

import "fmt"

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
}
