package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	freq := make(map[rune]int)
	for _, char := range angka {
		freq[char]++
	}
	var result []int
	for _, char := range angka {
		if freq[char] == 1 {
			result = append(result, int(char-'0'))
		}
	}
	return result
}

func main() {
	fmt.Println(munculsekali("1234123"))    // [4]
	fmt.Println(munculsekali("76523752"))   // [6, 3]
	fmt.Println(munculsekali("12345"))      // [ 1 2 3 4 5]
	fmt.Println(munculsekali("1122334455")) // []
	fmt.Println(munculsekali("0872504"))    // [8 7 2 5 4]
}
