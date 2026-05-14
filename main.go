package main

import "fmt"

func main() {
	var pilihan int
	fmt.Println("\n|===== CINE REVIEW =====|")
	fmt.Println("| 1. Tambah Film        |")
	fmt.Println("| 2. Tampilkan Film     |")
	fmt.Println("| 3. Urutkan Rating     |")
	fmt.Println("| 4. Urutkan Judul      |")
	fmt.Println("| 5. Cari Film          |")
	fmt.Println("| 6. Statistik Film     |")
	fmt.Println("|=======================|")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		tambahFilm()
	case 2:
		tampilFilm()
	case 3:
		selectionSortRating()
		tampilFilm()
	case 4:
		insertionSortJudul()
		tampilFilm()
	case 5:
		statistikFilm()
	case 6:
		var cari string
		fmt.Print("Masukkan judul film: ")
		fmt.Scan(&cari)
		idx := sequentialSearchJudul(cari)
		if idx == -1 {
			fmt.Println("Film tidak ditemukan.")
		} else {
			fmt.Printf("Film ditemukan pada indeks %d\n", idx)

		}
	case 0:
		return // atau break (tergantung konteks loop)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
