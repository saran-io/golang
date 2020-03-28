// For continued
// The init and post statements are optional.
// Tamil Translation

package main

import "fmt"

func main() {

	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println("For loop sum", sum)
}
