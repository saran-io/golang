// Appending to a slice
// It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
//
// func append(s []T, vs ...T) []T
// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
//
// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
//
// If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
//
// (To learn more about slices, read the Slices: usage and internals article.)

// Tamil translation

package main

import "fmt"

func main() {
	fmt.Print("\nAppend slice sample go program\n")
	var s []int
	printSlice(s)

	// append on nil slicces

	s = append(s, 0)
	printSlice(s)

	//the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//we can more element at a time
	s = append(s, 2, 3, 4, 5)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("Append slice len=%d cap=%d %v\n", len(s), cap(s), s)
}
