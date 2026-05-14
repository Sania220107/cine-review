package main

import (
	"fmt"
)

func statistikFilm() {

	if jumlahFilm == 0 {
		fmt.Println("Tidak ada data film.")
		return
	}

	genreCount := make(map[string]int)
	var totalRating float64 = 0.0

	for i := 0; i < jumlahFilm; i++ {
		genreCount[daftarFilm[i].genre]++
		totalRating += daftarFilm[i].rating
	}

	fmt.Println("\n=== STATISTIK FILM ===")
	fmt.Println("Jumlah film per genre:")

	for genre, count := range genreCount {
		fmt.Printf("  %-15s : %d film\n", genre, count)
	}

	averageRating := totalRating / float64(jumlahFilm)
	fmt.Printf("\nRata-rata rating seluruh koleksi : %.2f\n", averageRating)
}
