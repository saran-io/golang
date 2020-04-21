/***

** This is for the exercise for stringer in Golang
*TODO  -
*TODO -Exercise: Stringers
*TODO - Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
*TODO -For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

*! This is an important message for comments

 */

package main

import "fmt"

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
