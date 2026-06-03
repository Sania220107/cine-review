package main

import "fmt"

func tekanEnter() {
	fmt.Print("\nTekan ENTER untuk kembali ke menu...")
	fmt.Scanln()
	fmt.Scanln()
}

func main() {
	var pilihan int

	for {

		fmt.Println("\n============================")
		fmt.Println("===      CINE REVIEW     ===")
		fmt.Println("============================")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Tampilkan Film")
		fmt.Println("3. Edit Film")
		fmt.Println("4. Hapus Film")
		fmt.Println("5. Urutkan Rating")
		fmt.Println("6. Urutkan Tahun Rilis")
		fmt.Println("7. Statistik Film")
		fmt.Println("8. Cari Film Sequential")
		fmt.Println("9. Cari Film Binary")
		fmt.Println("0. Keluar")
		fmt.Println("============================")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahFilm()
			tekanEnter()

		case 2:
			tampilFilm()
			tekanEnter()

		case 3:
			editFilm()
			tekanEnter()

		case 4:
			hapusFilm()
			tekanEnter()

		case 5:
			selectionSortRating()
			fmt.Println("\nFilm berhasil diurutkan berdasarkan rating!")
			tampilFilm()
			tekanEnter()

		case 6:
			insertionSortTahun()
			fmt.Println("\nFilm berhasil diurutkan berdasarkan tahun rilis!")
			tampilFilm()
			tekanEnter()

		case 7:
			statistikFilm()
			tekanEnter()

		case 8:

			var pilihCari int
			var cari string

			fmt.Println("\n=== CARI FILM SEQUENTIAL ===")
			fmt.Println("1. Cari berdasarkan Judul")
			fmt.Println("2. Cari berdasarkan Genre")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilihCari)

			if pilihCari == 1 {

				fmt.Print("Masukkan judul film: ")
				fmt.Scan(&cari)

				idx := sequentialSearchJudul(cari)

				if idx == -1 {
					fmt.Println("Film tidak ditemukan.")
				} else {
					fmt.Printf("Film ditemukan pada indeks %d\n", idx)
				}

			} else if pilihCari == 2 {

				fmt.Print("Masukkan genre film: ")
				fmt.Scan(&cari)

				sequentialSearchGenre(cari)

			} else {
				fmt.Println("Pilihan tidak valid!")
			}

			tekanEnter()

		case 9:

			var cari string

			fmt.Println("\n=== CARI FILM BINARY ===")
			fmt.Print("Masukkan judul film: ")
			fmt.Scan(&cari)

			idx := binarySearchJudul(cari)

			if idx == -1 {
				fmt.Println("Film tidak ditemukan.")
			} else {
				fmt.Printf("Film ditemukan pada indeks %d\n", idx)
			}

			tekanEnter()

		case 0:
			fmt.Println("\nTerima kasih telah menggunakan Cine Review 🎬")
			return

		default:
			fmt.Println("\nPilihan tidak valid!")
			tekanEnter()
		}
	}
}