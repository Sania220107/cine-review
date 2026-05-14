package main

import "fmt"

func selectionSortRating() {
	n := len(daftarFilm)
	if n == 0 {
		fmt.Println("Daftar film kosong, tidak ada yang bisa diurutkan.")
		return
	}

	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if daftarFilm[j].rating > daftarFilm[idxMax].rating {
				idxMax = j
			}
		}

		daftarFilm[i], daftarFilm[idxMax] = daftarFilm[idxMax], daftarFilm[i]
	}
}

func insertionSortJudul() {
	n := len(daftarFilm)
	if n == 0 {
		fmt.Println("Daftar film kosong, tidak ada yang bisa diurutkan.")
		return
	}

	for i := 1; i < n; i++ {
		key := daftarFilm[i]
		j := i - 1

		for j >= 0 && daftarFilm[j].judul > key.judul {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
}