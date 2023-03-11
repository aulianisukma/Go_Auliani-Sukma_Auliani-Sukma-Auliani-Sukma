package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	// goroutine untuk mengisi channel dengan bilangan kelipatan 3
	go func() {
		for i := 0; i < 20; i++ {
			if i%3 == 0 {
				ch <- i
			}
		}
		close(ch)
	}()

	// membaca isi channel dan mencetak bilangan kelipatan 3
	for v := range ch {
		fmt.Println(v)
	}
}