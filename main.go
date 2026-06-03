package main

import "fmt"

func main() {
	var pilihan int

	for {

		fmt.Println("\n=================================")
		fmt.Println("         🎬 CINE REVIEW          ")
		fmt.Println("=================================")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Tampilkan Film")
		fmt.Println("3. Urutkan Rating")
		fmt.Println("4. Urutkan Judul")
		fmt.Println("5. Statistik Film")
		fmt.Println("6. Cari Film")
		fmt.Println("0. Keluar")
		fmt.Println("=================================")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahFilm()

		case 2:
			tampilFilm()

		case 3:
			selectionSortRating()
			fmt.Println("\nFilm berhasil diurutkan berdasarkan rating!")
			tampilFilm()

		case 4:
			insertionSortJudul()
			fmt.Println("\nFilm berhasil diurutkan berdasarkan judul!")
			tampilFilm()

		case 5:
			tampilFilm()
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
			fmt.Println("\nTerima kasih telah menggunakan Cine Review 🎬")
			return

		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}