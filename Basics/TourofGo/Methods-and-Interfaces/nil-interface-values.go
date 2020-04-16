/**
* Beginning of the comment block

* * this is for the nil interface value setup
* ! this is related to the interface section
* TODO-1  Nil interface values A nil interface value holds neither value nor concrete type.
* TODO-2  Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

* ? Becareful with the methods you declare. This is very important for Golang


* End of the comment block. All of the programming logics would be evaulated later.
 */

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
