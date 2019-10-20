package main

import "fmt"

func main() {
	favouriteThings := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
		`ketchum_ash`:     {`catching pokemon`, `pikachu`, `pokemon tournaments`},
	}
	if v, ok := favouriteThings[`ketchum_ash`]; ok {
		fmt.Println(v)
		delete(favouriteThings, `ketchum_ash`)
	}

	for key, value := range favouriteThings {
		fmt.Printf("Name: %s\n", key)
		for index, thing := range value {
			fmt.Printf("Index: %d, Favourite thing: %s\n", index, thing)
		}
	}
}
