package main

import (
	"fmt"
)

func pritnSquare(rows int, col int) {
	for i := 0; i <= rows; i++ {
		for j := 0; j <= col; j++ {
			if i == 0 || i == rows || j == 0 || j == col {
				fmt.Println("*")

			} else {
				fmt.Println(" ")
			}
		}

	}

}
func main() {
	fmt.Println("enter the number of rows")
	var rows int
	fmt.Scanln(&rows)
	fmt.Println("enter the number of coloumns")
	var col int
	fmt.Scanln(&col)
	fmt.Println(pritnSquare(rows, col))

}
