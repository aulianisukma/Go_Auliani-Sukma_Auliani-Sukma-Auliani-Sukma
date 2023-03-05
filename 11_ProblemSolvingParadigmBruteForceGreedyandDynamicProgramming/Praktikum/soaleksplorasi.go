package main

import "fmt"

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			roman += romans[i]
		}
	}

	return roman
}

func main() {
	fmt.Println(intToRoman(4))     // IV
	fmt.Println(intToRoman(9))     // IX
	fmt.Println(intToRoman(23))    // XXIII
	fmt.Println(intToRoman(2021))  // MMXXI
	fmt.Println(intToRoman(1646))  // MDCXLVI
}
