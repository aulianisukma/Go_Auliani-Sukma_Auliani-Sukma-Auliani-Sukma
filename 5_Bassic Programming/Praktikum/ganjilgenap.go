package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scanln(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Grade Anda adalah A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("Grade Anda adalah B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("Grade Anda adalah C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("Grade Anda adalah D")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("Grade Anda adalah E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
