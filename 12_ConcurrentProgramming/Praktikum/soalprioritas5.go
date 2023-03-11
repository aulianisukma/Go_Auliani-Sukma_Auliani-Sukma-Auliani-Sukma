package main

import (
	"fmt"
	"sync"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	result := make(map[rune]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, ch := range text {
		wg.Add(1)
		go func(ch rune) {
			mutex.Lock()
			result[ch]++
			mutex.Unlock()
			wg.Done()
		}(ch)
	}
	
	wg.Wait()

	for ch, count := range result {
		if count > 0 && ch != ' ' && ch != ',' && ch != '.' {
			fmt.Printf("%c: %d\n", ch, count)
		}
	}
}