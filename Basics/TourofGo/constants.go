// Constants
// Constants are declared like variables, but with the const keyword.
//
// Constants can be character, string, boolean, or numeric values.
//
// Constants cannot be declared using the := syntax.

// Tamil translation

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Tamil"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
