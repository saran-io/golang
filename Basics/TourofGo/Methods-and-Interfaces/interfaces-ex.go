// Go language interfaces are different from other languages. In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.

// Interaces in Golang is like set of methods to carry out specific task.
// For example learning would be an interface which would have diff methods of learning
// available. Like online, languages etc
//
// In Golang all the package Methods and Interfaces

//implementing interfaces.

package main

import "fmt"

//create the interace learning
//it will have the methods to be created for learning
type learning interface {

	//methods declaration
	Mode() string
	Fees() float64
	Coaches() string
}

type feeder struct {
	mode_of_class string
	fee_amount    float64
	coach_name    string
}

//implement the methods for the interfaces

func (f feeder) Mode() string {
	fmt.Print("Mode of learning platform ")
	return f.mode_of_class
}

func (f feeder) Fees() float64 {
	fmt.Print("Fees ")
	return f.fee_amount
}

func (f feeder) Coaches() string {
	fmt.Print("Coach for the Platform ")
	return f.coach_name
}

func main() {
	var l learning
	l = feeder{"Online", 56, "John"}
	fmt.Println(l.Mode())
	fmt.Println(l.Fees())
	fmt.Println(l.Coaches())
}
