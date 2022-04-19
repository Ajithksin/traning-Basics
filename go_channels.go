package main

import (
	"fmt"
	"time"
)

func send(c chan int) {
	c <- 5
}
func send_V_again(st []int, ch chan []int) {
	var sr []int
	for _, v := range st {
		sr = append(sr, v)
		fmt.Println("yes")
	}
	ch <- sr

}
func main() {
	fmt.Println("go Channnels start")
	st := []int{1, 2, 3, 4}
	val2 := make(chan []int)
	go send_V_again(st, val2)
	//va := <-val2
	//fmt.Println(va)
	time.Sleep(time.Second)
	values := make(chan int)
	go send(values)
	value := <-values
	fmt.Println(value)
}
