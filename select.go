package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go ajit(ch1)

	go kumar(ch2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}

	}
	//select
}
func ajit(ch chan string) {
	ch <- " woooooo "
}
func kumar(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- " kumarrrrr"
}
