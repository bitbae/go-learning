package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range mySlice {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	fmt.Printf("%T", mySlice)
}
