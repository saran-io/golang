/**

** Begin: Comment block

*TODO - Implement Error type sample for the Golang.
*! Definition: Its a built in interface in Golang like stringer. When you implement the program
*! looks for the value and implement it run time.package Errors


Excerpts from official Golang
Errors
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.

** End: Comment block

*/

package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	When time.Time
	What string
}

func (e *Myerror) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &Myerror{
		time.Now(), "it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
