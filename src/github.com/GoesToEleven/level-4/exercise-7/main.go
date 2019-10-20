package main

import "fmt"

func main() {
	slice1 := [][]string{}
	slice2 := []string{"James", "Bond", "Shaken, not stirred"}
	slice3 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	slice1 = append(slice1, slice2, slice3)

	for _, value := range slice1 {
		fmt.Println("Value: ", value)
		for _, value1 := range value {
			fmt.Println("Data: ", value1)
		}
	}
}
