/**

** Check the Stringer Type


* TODO - Stringers
* TODO - One of the most ubiquitous interfaces is Stringer defined by the fmt package.
* TODO -
* TODO - type Stringer interface {
* TODO -     String() string
* TODO - }
* TODO - A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

* ! Stringer is a interface in golang

 */

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v  (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Samuel Jackson\n", 49}
	z := Person{"Mary Poppins\n", 89}
	fmt.Println(a, z)
}
