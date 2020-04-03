// Methods continued
// You can declare a method on non-struct types, too.
//
// In this example we see a numeric type MyFloat with an Abs method.
//
// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

// Tamil Translation

package main

import (
	"fmt"
	"math"
)

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat64(-math.Sqrt2)
	fmt.Println(f.Abs())
}
