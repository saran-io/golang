/**
** Begin: Comment block

* TODO - Type assertions
* TODO - A type assertion provides access to an interface value's underlying concrete value.
* TODO - t := i.(T)
* TODO - This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
* TODO - If i does not hold a T, the statement will trigger a panic.
* TODO -
* TODO - To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
* TODO -
* TODO - t, ok := i.(T)
* TODO - If i holds a T, then t will be the underlying value and ok will be true.
* TODO -
* TODO - If not, ok will be false and t will be the zero value of type T, and no panic occurs.
* TODO -
* TODO - Note the similarity between this syntax and that of reading from a map.

* ! VALIDATE THE CODE BLOCK BEFORE YOU EXECUTE THE PROGRAM


** End: Comment block
 */

package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
