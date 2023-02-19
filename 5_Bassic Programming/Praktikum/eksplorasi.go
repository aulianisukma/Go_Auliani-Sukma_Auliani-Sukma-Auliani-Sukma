package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("aya")
	fmt.Scanln(&input)
	if isPalindrome(input) {
		fmt.Println(input, "aya adalah palindrome.")
	} else {
		fmt.Println(input, "bukan palindrome.")
	}
}
