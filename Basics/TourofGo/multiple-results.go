package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Multiple results test")
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
