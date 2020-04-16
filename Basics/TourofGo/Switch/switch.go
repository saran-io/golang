/**

** Switch Statement in Go
** A switch statement is a multiway branch statement. It provides an efficient way to transfer the execution to different parts of a code based on the value(also called case) of the expression. Go language supports two types of switch statements:

** 1. Expression Switch
** 2. Type Switch


*! the examples given here few samples in switch.
 */

package main

import "fmt"

// example-1 . This is for the day validation using expression switch

func switchex() {
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}
}

func switchex2() {

	fmt.Println("Second example of swtich in golang")
	var value int = 2

	switch {

	case value == 1:
		fmt.Println("Hello Golang. Whats up buddy")

	case value == 2:
		fmt.Println("Is everything alrite! I just heard a boom")

	case value == 2:
		fmt.Println("May the force be with you")
	}
}

func switchex3() {
	fmt.Println("This is the best part of Golang. Switch can validate multiple values")

	var value string = "five"

	switch value {
	case "one":
		fmt.Println("C language")
	case "two":
		fmt.Println("C++ language")
	case "three", "five", "four":
		fmt.Println("Golang - The master")

	}
}

func main() {
	switchex()
	switchex2()
	switchex3()
}
