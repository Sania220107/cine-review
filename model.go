package main

type Film struct {
	judul      string
	genre      string
	tahun      int
	rating     float64
	deskripsi  string
}

var daftarFilm [100]Film
var jumlahFilm int