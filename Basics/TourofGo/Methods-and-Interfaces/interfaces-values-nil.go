// Interface values with nil underlying values
// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
//
// In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
//
// Note that an interface value that holds a nil concrete value is itself non-nil.

/**

This is the comments section for the given golan program
* * Important information about the given program.
* ! careful with the deprecated methods
* TODO-1 : Follow the correct methods for implmenting the logicals change
* TODO-2 : Create the interface
* TODO-3 : Declare the methods
* TODO-4 : Create the main function and write it down
* ? Be Careful with the method explaining the methods
* @param the parameter method for this one

*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {

	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}
