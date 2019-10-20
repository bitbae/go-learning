package main

import "fmt"

func main() {
	favouriteThings := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	favouriteThings[`ketchum_ash`] = []string{`catching pokemon`, `pikachu`, `pokemon tournaments`}
	for key, value := range favouriteThings {
		fmt.Printf("Name: %s\n", key)
		for index, thing := range value {
			fmt.Printf("Index: %d, Favourite thing: %s\n", index, thing)
		}
	}
}
