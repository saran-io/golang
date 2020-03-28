// Pointers
// Go has pointers. A pointer holds the memory address of a value.
//
// The type *T is a pointer to a T value. Its zero value is nil.
//
// var p *int
// The & operator generates a pointer to its operand.
//
// i := 42
// p = &i
// The * operator denotes the pointer's underlying value.
//
// fmt.Println(*p) // read i through the pointer p
// p = 21         // set i through the pointer p
// This is known as "dereferencing" or "indirecting".
//
// Unlike C, Go has no pointer arithmetic.

package main

import "fmt"

func main() {
	//declare the variables
	i, j := 42, 2701                          //declaration
	p := &i                                   //set the operatnd
	fmt.Println("Pointer value", p, *p, i, j) /// get the value of the value via the operand
	*p = 21                                   //set hhe value if I via pointer/operand
	fmt.Println("Value of the i reassgiend via pointer", i)

	p = &j // point to j
	*p = *p / 37
	fmt.Println("New value of J", j)
}
