/**
 ** beginnig of the comment block

 *! This is for the empty interfaces
*TODO-1:
*TODO-1: The empty interface
*TODO-1: The interface type that specifies zero methods is known as the empty interface:
*TODO-1: interface{}
*TODO-1: An empty interface may hold values of any type. (Every type implements at least zero methods.)
*TODO-1:  Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

*? Check other updates later in this code block

** End of the commenting block
*/

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
