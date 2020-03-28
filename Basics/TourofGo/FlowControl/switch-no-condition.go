//
// Switch with no condition
// Switch without a condition is the same as switch true.
//
// This construct can be a clean way to write long if-then-else chains.
//

// Tamil Translation

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Current Time value is", t)
	thour := t.Hour()
	fmt.Println("Current Time value is", thour)

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
