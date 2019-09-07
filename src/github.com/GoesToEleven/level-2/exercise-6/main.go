package main

import "fmt"

const (
	thisyear = 2019
	a        = thisyear + iota
	b        = thisyear + iota
	c        = thisyear + iota
	d        = thisyear + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
