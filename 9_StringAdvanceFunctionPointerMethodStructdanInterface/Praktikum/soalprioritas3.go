package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a) {
		return a
	} else {
		return ""
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
