// Variables with initializers
// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

// Tamil Translation

package main

import "fmt"

var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Print("Varaibles with initializers in Go\n")
	fmt.Println(i, j, c, python, java)
}
