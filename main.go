package main

import "fmt"

func main() {
	var pilihan int
	fmt.Println("\n|===== CINE REVIEW =====|")
	fmt.Println("| 1. Tambah Film        |")
	fmt.Println("| 2. Tampilkan Film     |")
	fmt.Println("| 3. Urutkan Rating     |")
	fmt.Println("| 4. Urutkan Judul      |")
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
	case 0:
		return // atau break (tergantung konteks loop)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}