package main

import "fmt"

func main() {
	menu := map[string]float64{
		"tea":    2,
		"samosa": 10,
	}
	fmt.Println(menu)
	fmt.Println(menu["tea"])

	//looping
	for k, v := range menu {
		fmt.Println(k, "--", v)
	}
	//updating
	menu["tea"] = 5
	fmt.Println(menu)
}
