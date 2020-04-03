// Variables
// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level. We see both in this example.

// Tamil Translation

package main

import "fmt"

var age int
var name, dept string
var promoted bool

func main() {
	age := 23
	name := "Rockstar"
	dept := "Engineering"
	promoted := "Not this time! Dude. "
	fmt.Print("Variables example go\n")
	fmt.Println(name, age, dept, promoted)
}
