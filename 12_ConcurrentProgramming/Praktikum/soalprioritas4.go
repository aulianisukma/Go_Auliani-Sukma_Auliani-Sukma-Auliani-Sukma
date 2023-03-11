package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	evenCount := 0
	oddCount := 0

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if num%2 == 0 {
				evenCount++
			} else {
				oddCount++
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total bilangan genap: %d\n", evenCount)
	fmt.Printf("Total bilangan ganjil: %d\n", oddCount)
}