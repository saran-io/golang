// Slice literals
// A slice literal is like an array literal without the length.
//
// This is an array literal:
//
// [3]bool{true, true, false}
// And this creates the same array as above, then builds a slice that references it:
//
// []bool{true, true, false}
// You are not allowed to mention the size of it. But you can add or modify it.
// Tamil translation

package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("slice literal test\n", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("Bool slice literals\n", r)

	//combine both and create a new slice literal
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Test array", s)
}
