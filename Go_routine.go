package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ";", i)
	}
}
func main() {
	f("direct")
	tm := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", tm.Format("2006-01-02 15:04:05.000000"))
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(1 * time.Minute)
	fmt.Println("done")
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", time.Now().Format("2006-01-02 15:04:05.000000"))
}
