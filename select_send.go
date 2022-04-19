package main

import "fmt"

func one(ch chan string) {
	ch <- "don don don "
}
func two(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go one(ch1)
	go two(ch2)
	for i := 0; i < 2; i++ {
		select {

		case msg1 := <-ch1:
			fmt.Println(msg1)
		case ch2 <- "to goTwo goroutine":
		}
	}
}
