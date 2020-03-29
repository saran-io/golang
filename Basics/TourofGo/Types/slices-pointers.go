// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying array.
//
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
//
// Other slices that share the same underlying array will see those changes.
// syntax: array [ low : high]
// you could store similar type in array

// Tamil Translation

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("Names array", names)

	//slice the array examples
	a := names[0:2]
	b := names[2:4]

	fmt.Println("sliced array\n", a, "\n", b)

	//modify the array
	b[0] = "xxx"

	fmt.Println("Modified array\n", a, "\n", b)
	fmt.Println("Modified array", names)

}
