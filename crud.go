package main

import "fmt"

func tambahFilm() {
	var film Film

	fmt.Println("\n=== TAMBAH FILM ===")

	fmt.Print("Judul      : ")
	fmt.Scan(&film.judul)

	fmt.Print("Genre      : ")
	fmt.Scan(&film.genre)

	fmt.Print("Tahun      : ")
	fmt.Scan(&film.tahun)

	fmt.Print("Rating     : ")
	fmt.Scan(&film.rating)

	fmt.Print("Deskripsi  : ")
	fmt.Scan(&film.deskripsi)

	daftarFilm = append(daftarFilm, film)

	fmt.Println("\nFilm berhasil ditambahkan!")
}

func tampilFilm() {

	fmt.Println("\n=== DAFTAR FILM ===")

	if len(daftarFilm) == 0 {
		fmt.Println("Belum ada data film.")
		return
	}

	for i, film := range daftarFilm {

		fmt.Printf("\nFilm ke-%d\n", i+1)
		fmt.Println("Judul     :", film.judul)
		fmt.Println("Genre     :", film.genre)
		fmt.Println("Tahun     :", film.tahun)
		fmt.Println("Rating    :", film.rating)
		fmt.Println("Deskripsi :", film.deskripsi)
	}
}

func editFilm() {

	var index int

	fmt.Println("\n=== EDIT FILM ===")

	if len(daftarFilm) == 0 {
		fmt.Println("Belum ada data film.")
		return
	}

	tampilFilm()

	fmt.Print("\nMasukkan nomor film yang ingin diedit: ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= len(daftarFilm) {
		fmt.Println("Nomor film tidak valid!")
		return
	}

	fmt.Print("Judul baru      : ")
	fmt.Scan(&daftarFilm[index].judul)

	fmt.Print("Genre baru      : ")
	fmt.Scan(&daftarFilm[index].genre)

	fmt.Print("Tahun baru      : ")
	fmt.Scan(&daftarFilm[index].tahun)

	fmt.Print("Rating baru     : ")
	fmt.Scan(&daftarFilm[index].rating)

	fmt.Print("Deskripsi baru  : ")
	fmt.Scan(&daftarFilm[index].deskripsi)

	fmt.Println("\nFilm berhasil diedit!")
}

func hapusFilm() {

	var index int

	fmt.Println("\n=== HAPUS FILM ===")

	if len(daftarFilm) == 0 {
		fmt.Println("Belum ada data film.")
		return
	}

	tampilFilm()

	fmt.Print("\nMasukkan nomor film yang ingin dihapus: ")
	fmt.Scan(&index)

	index--

	if index < 0 || index >= len(daftarFilm) {
		fmt.Println("Nomor film tidak valid!")
		return
	}

	daftarFilm = append(daftarFilm[:index], daftarFilm[index+1:]...)

	fmt.Println("\nFilm berhasil dihapus!")
}