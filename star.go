package main

import (
	"fmt"
)

func star(number int) int {
	if number < 0 {
		return 0
	} else {
		fmt.Println("8")
		return star(number - 1)
	}

}
func main() {
	fmt.Println(star(5))
}
