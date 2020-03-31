// Slices of slices
// Slices can contain any type, including other slices.

// Tamil Translation

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nSlice of slices test\n")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//the player takes turns
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"
	fmt.Println("board", board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
