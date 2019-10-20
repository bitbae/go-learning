package main

import "fmt"

func main() {
	var myArray [5]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	for index, value := range myArray {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	fmt.Printf("%T", myArray)
}
