package main

import "fmt"

func statistikFilm() {

	if len(daftarFilm) == 0 {
		fmt.Println("\nTidak ada data film.")
		return
	}

	genreCount := make(map[string]int)
	var totalRating float64

	for i := 0; i < len(daftarFilm); i++ {

		genreCount[daftarFilm[i].genre]++
		totalRating += daftarFilm[i].rating
	}

	fmt.Println("\n=== STATISTIK FILM ===")
	fmt.Println("Jumlah film per genre:")

	for genre, jumlah := range genreCount {
		fmt.Printf("%s : %d film\n", genre, jumlah)
	}

	rataRata := totalRating / float64(len(daftarFilm))

	fmt.Printf("\nRata-rata rating seluruh koleksi : %.2f\n", rataRata)
}