package main

import (
	"fmt"
)

func pattern(num int, i int) int {

	if num < 1 {
		return 0
	}
	if i <= num {
		fmt.Println("*")
		return pattern(num, i+1)
	} else {
		fmt.Println()
		return pattern(num-1, 1)
	}
}
func main() {
	fmt.Println("enter the number")
	var num int
	fmt.Scanln(&num)
	pattern(num, 1)
}
