package main

import "fmt"

func main() {
	x := 1
	if x == 1 {
		fmt.Println("x is one")
	} else if x == 2 {
		fmt.Println("x is two")
	} else {
		fmt.Println("x is neither one nor two")
	}
}
