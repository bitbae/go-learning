package main

import "fmt"

// Declare variable with identifier 'z' and type int
// Assigns the zero value of type int to 'z'
var z int

// Declare and assign = initialization
var y = 43

func main() {
	// Short declaration operator
	// Declare a variable and assign a value of a certain type at the same time
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
