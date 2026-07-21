package main

import "fmt"

func main() {
	ShowGoroutine(1)
}

func ShowGoroutine(id int) {
	fmt.Print("Goroutine #%d\n", id)

}
