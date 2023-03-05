package main

import "fmt"

func Frog(jumps []int) int {
	maxDiff := 0
	for i := 1; i < len(jumps); i++ {
		diff := jumps[i] - jumps[i-1]
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))    // 30
	fmt.Println(Frog([]int{30, 60, 10, 60, 50})) // 40
}
