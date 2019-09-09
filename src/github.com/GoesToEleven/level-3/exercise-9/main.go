package main

import "fmt"

func main() {
	favSport := "frisbee"
	switch favSport {
	case "frisbee":
		fmt.Println(favSport)
	case "badminton":
		fmt.Println("baddy")
	default:
		fmt.Println("none")
	}
}
