package main

import "fmt"

func tambahFilm() {
	if jumlahFilm < 100 {
		var f Film
		fmt.Print("Judul: ")
		fmt.Scan(&f.judul)
		fmt.Print("Genre: ")
		fmt.Scan(&f.genre)
		fmt.Print("Tahun: ")
		fmt.Scan(&f.tahun)
		fmt.Print("Rating (0.0 - 10.0): ")
		fmt.Scan(&f.rating)
		fmt.Print("Deskripsi: ")
		fmt.Scan(&f.deskripsi)

		daftarFilm[jumlahFilm] = f
		jumlahFilm++
		fmt.Println("Film berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func tampilFilm() {
	fmt.Println("\n--- Daftar Film ---")
	for i := 0; i < jumlahFilm; i++ {
		f := daftarFilm[i]
		fmt.Printf("%d. [%s] Genre: %s, Tahun: %d, Rating: %.1f\n", i+1, f.judul, f.genre, f.tahun, f.rating)
	}
}