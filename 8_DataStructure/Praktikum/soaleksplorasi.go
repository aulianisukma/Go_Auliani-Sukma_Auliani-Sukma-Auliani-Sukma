package main

import "fmt"

func main() {
	// Definisikan matriks persegi sebagai array 2D
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// Hitung jumlah diagonal kiri ke kanan
	var sumLR int
	for i := 0; i < len(matrix); i++ {
		sumLR += matrix[i][i]
	}

	// Hitung jumlah diagonal kanan ke kiri
	var sumRL int
	for i := 0; i < len(matrix); i++ {
		sumRL += matrix[i][len(matrix)-1-i]
	}

	// Hitung selisih absolut antara jumlah diagonal kiri ke kanan dan diagonal kanan ke kiri
	diff := sumLR - sumRL
	if diff < 0 {
		diff = -diff
	}

	// Tampilkan hasil
	fmt.Printf("Matriks: %v\n", matrix)
	fmt.Printf("Selisih absolut antara jumlah diagonalnya: %d\n", diff)
}
