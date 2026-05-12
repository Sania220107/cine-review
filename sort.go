package main
import "fmt"

// mengurutkan rating dari tertinggi
func selectionSortRating() {
	if jumlahFilm == 0 {
		fmt.Print("Daftar Film Kosong")
		return
	}

	for i := 0; i < jumlahFilm - 1; i++ {
		idxMax := i
		for j := i + 1; j < jumlahFilm; j++ {
			if daftarFilm[j].rating > daftarFilm[idxMax].rating {
				idxMax = j
			}
		}
		temp := daftarFilm[idxMax]
		daftarFilm[idxMax] = daftarFilm[i]
		daftarFilm[i] = temp
	}
	fmt.Println("Berhasil mengurutkan film berdasarkan rating  tertinggi")
}

// mengurutkan judul sesuai abjad
func insertionSortJudul() {
	if jumlahFilm == 0 {
		fmt.Println("Daftar film kosong.")
		return
	}
	for i := 1; i < jumlahFilm; i++ {
		key := daftarFilm[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && daftarFilm[j].judul > key.judul {
			daftarFilm[j+1] = daftarFilm[j]
			j--
		}
		daftarFilm[j+1] = key
	}
	fmt.Println("Berhasil mengurutkan film berdasarkan judul.")
}