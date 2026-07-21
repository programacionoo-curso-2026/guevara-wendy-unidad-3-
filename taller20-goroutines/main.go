package main

import (
	"fmt"
	"time"
)

func main() {
	go ShowGoroutine(1)
	time.Sleep(10 * time.Second)

}

func ShowGoroutine(id int) {
	fmt.Print("Goroutine #%d\n", id)

}
