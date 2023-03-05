package main

import (
	"fmt"
)

func primeX(number int) int {
	count := 0 // variabel untuk menghitung bilangan prima
	prime := 2 // variabel untuk menyimpan bilangan prima

	for count < number {
		isPrime := true // variabel untuk menandai apakah bilangan saat ini prima atau tidak

		// memeriksa apakah bilangan saat ini prima
		for i := 2; i < prime; i++ {
			if prime%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++ // jika bilangan saat ini prima, tambahkan ke hitungan
		}

		prime++ // lanjut ke bilangan berikutnya
	}

	return prime - 1 // mengembalikan bilangan prima terakhir yang ditemukan
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) //29
}
