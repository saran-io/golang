//functions go file for the sample programming

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Print("Functions program for addition\n")
	fmt.Println(add(10, 23))
}
