package main

import "fmt"

func sequentialSearchJudul(judul string) int {

	for i := 0; i < len(daftarFilm); i++ {

		if daftarFilm[i].judul == judul {
			return i
		}
	}

	return -1
}

func sequentialSearchGenre(genre string) {

	found := false

	fmt.Println("\n=== HASIL PENCARIAN GENRE ===")

	for i := 0; i < len(daftarFilm); i++ {

		if daftarFilm[i].genre == genre {

			fmt.Printf("\nFilm ke-%d\n", i+1)
			fmt.Println("Judul     :", daftarFilm[i].judul)
			fmt.Println("Genre     :", daftarFilm[i].genre)
			fmt.Println("Tahun     :", daftarFilm[i].tahun)
			fmt.Println("Rating    :", daftarFilm[i].rating)
			fmt.Println("Deskripsi :", daftarFilm[i].deskripsi)

			found = true
		}
	}

	if !found {
		fmt.Println("Film dengan genre tersebut tidak ditemukan.")
	}
}

func binarySearchJudul(judul string) int {

	if len(daftarFilm) == 0 {
		return -1
	}

	type item struct {
		idx   int
		judul string
	}

	temp := make([]item, len(daftarFilm))

	for i := 0; i < len(daftarFilm); i++ {

		temp[i] = item{
			idx:   i,
			judul: daftarFilm[i].judul,
		}
	}

	for i := 1; i < len(temp); i++ {

		key := temp[i]
		j := i - 1

		for j >= 0 && temp[j].judul > key.judul {
			temp[j+1] = temp[j]
			j--
		}

		temp[j+1] = key
	}

	low := 0
	high := len(temp) - 1

	for low <= high {

		mid := (low + high) / 2

		if temp[mid].judul == judul {
			return temp[mid].idx
		} else if temp[mid].judul < judul {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}