// Zero values
// Variables declared without an explicit initial value are given their zero value.
//
// The zero value is:
//
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.\
// Check the format in packages in go for more details : https://golang.org/pkg/fmt/

// Tamil Translation

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Values\n %v\n %v\n %v\n %q\n", i, f, b, s)
}
