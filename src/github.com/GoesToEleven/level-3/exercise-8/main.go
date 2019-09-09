package main

import "fmt"

func main() {
	x := 1
	switch {
	case x == 1:
		fmt.Println("x is one")
	case x == 2:
		fmt.Println("x is two")
	}
}
