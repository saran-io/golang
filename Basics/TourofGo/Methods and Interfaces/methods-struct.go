// In Go language, you are allowed to define a method whose receiver is of a struct type. This receiver is accessible inside the method as shown in the below example:
// Go program to ilustrate teh method with struct type received

// package Methods and Interfaces

package main

import "fmt"

//Author structure
type author struct {
	name     string
	branch   string
	articles int
	salary   int
}

//Method with a receiver of author type

func (a author) show() {

	fmt.Println("Author's name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published Articles: ", a.articles)
	fmt.Println("Salary: ", a.salary)
}

// main function

func main() {

	//initalizing the values of author structured

	res := author{
		name:     "Mike",
		branch:   "Ecology",
		articles: 203,
		salary:   34000,
	}

	res.show()
}
