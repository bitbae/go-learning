package main

import "fmt"

func main() {
	date := 1995
	for {
		if date == 2020 {
			break
		}
		fmt.Println(date)
		date++
	}
}
