// Struct Fields
// Struct fields are accessed using a dot.
// Tamil Translation

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{23, 76}
	v.X = 4
	fmt.Println(v.X)
}
