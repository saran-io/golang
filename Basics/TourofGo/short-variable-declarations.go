// Short variable declarations
// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

// Tamil Translation

package main

import "fmt"

func main() {
	fmt.Print("Declatre the variables \n")
	var i, j int = 4, 5
	x := 12

	c, python, java := true, false, "no!"
	fmt.Print("Result of the variables declaration\n")
	fmt.Println(i, j, x, c, python, java)
	fmt.Println("Program completed")
}
