package main

import "fmt"

func chans(ch chan int) {
	fmt.Println(200 + <-ch)
}
func main() {
	fmt.Println("start main method")

	ch := make(chan int)

	aa := 20 + <-ch

	go chans(ch)
	ch <- 23
	fmt.Println("end main method")
	fmt.Println(aa)

}
