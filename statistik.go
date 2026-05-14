package main

import (
	"fmt"
)

func statistikFilm() {
	if len(daftarFilm) == 0 {
		fmt.Println("Tidak ada data film.")
		return
	}

	genreCount := make(map[string]int)
	totalRating := 0.0

	for _, film := range daftarFilm {
		genreCount[film.genre]++
		totalRating += film.rating
	}

	fmt.Println("\n=== STATISTIK FILM ===")
	fmt.Println("Jumlah film per genre:")
	for genre, count := range genreCount {
		fmt.Printf("  %s : %d film\n", genre, count)
	}

	averageRating := totalRating / float64(len(daftarFilm))
	fmt.Printf("\nRata-rata rating seluruh koleksi: %.2f\n", averageRating)
}
