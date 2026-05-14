package main

func sequentialSearchJudul(judul string) int {
	for i := 0; i < jumlahFilm; i++ {
		if daftarFilm[i].judul == judul {
			return i
		}
	}
	return -1
}

func binarySearchJudul(judul string) int {
	if jumlahFilm == 0 {
		return -1
	}

	type item struct {
		idx   int
		judul string
	}

	temp := make([]item, jumlahFilm)
	for i := 0; i < jumlahFilm; i++ {
		temp[i] = item{idx: i, judul: daftarFilm[i].judul}
	}

	for i := 1; i < jumlahFilm; i++ {
		key := temp[i]
		j := i - 1
		for j >= 0 && temp[j].judul > key.judul {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = key
	}

	low, high := 0, jumlahFilm-1
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
