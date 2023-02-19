package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	// Minta input sisi atas, sisi bawah, dan tinggi dari pengguna
	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&sisiAtas)
	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&sisiBawah)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	// Hitung luas trapesium
	luas := 0.5 * (sisiAtas + sisiBawah) * tinggi

	// Tampilkan hasil
	fmt.Println("Luas trapesium adalah", luas)
}
