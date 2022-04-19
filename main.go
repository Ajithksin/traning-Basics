package main

import (
	"fmt"
	"time"
)

func main() {
	go count("ajit")
	count("kumar")

}

func count(things string) {
	for i := 0; true; i++ {
		fmt.Println(i, things)
		time.Sleep(time.Second + 1)

	}
}
